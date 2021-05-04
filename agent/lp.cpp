#include "lp.hpp"

#include <memory>
#include <sstream>
#include <locale>
#include <codecvt>

#include <windows.h>
#include <winhttp.h>

ListeningPost::ListeningPost(const char *hostname, int port, const char * userAgent) 
{
    this->hostname = wstring_convert<codecvt_utf8_utf16<wchar_t>>().from_bytes(hostname);
    this->port = port;
    this->userAgent = wstring_convert<codecvt_utf8_utf16<wchar_t>>().from_bytes(userAgent);

}

string ListeningPost::getData()
{
    DWORD size = 0;
    DWORD downloaded = 0;
    unique_ptr<char[]> buffer;
    stringstream result;

    HINTERNET hSession = WinHttpOpen(this->userAgent.c_str(), WINHTTP_ACCESS_TYPE_DEFAULT_PROXY, WINHTTP_NO_PROXY_NAME, WINHTTP_NO_PROXY_BYPASS, 0);
    HINTERNET hConnect = WinHttpConnect(hSession, this->hostname.c_str(), this->port, 0);
    HINTERNET hRequest = WinHttpOpenRequest(hConnect, L"GET", NULL, NULL, NULL, NULL, 0);

    if (!WinHttpSendRequest(hRequest, WINHTTP_NO_ADDITIONAL_HEADERS, 0, WINHTTP_NO_REQUEST_DATA, 0, 0, 0))
        throw runtime_error("Error sending request");

    if (!WinHttpReceiveResponse(hRequest, NULL)) 
        throw runtime_error("Error receiving response");

    do
    {
        if (!WinHttpQueryDataAvailable(hRequest, &size))
            throw runtime_error("Error getting data available");


        buffer = make_unique<char[]>(size);

        if (!WinHttpReadData(hRequest, (LPVOID)buffer.get(), size, &downloaded)) {
            throw runtime_error("Error reading data");
        }

        result.write(buffer.get(), downloaded);

    } while (size > 0);

    return result.str();
}