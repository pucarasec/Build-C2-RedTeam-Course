#include <stdio.h>
#include <windows.h>
#include <winhttp.h>
//#include <wininet.h>

int main(int argc, char **argv)
{
    DWORD dwSize = 0;
    DWORD dwDownloaded = 0;
    LPSTR pszOutBuffer;

    HINTERNET hSession = WinHttpOpen(L"curl/7.64.1", WINHTTP_ACCESS_TYPE_DEFAULT_PROXY, WINHTTP_NO_PROXY_NAME, WINHTTP_NO_PROXY_BYPASS, 0);
    HINTERNET hConnect = WinHttpConnect(hSession, L"ipinfo.io", INTERNET_DEFAULT_HTTP_PORT, 0);
    HINTERNET hRequest = WinHttpOpenRequest(hConnect, L"GET", NULL, NULL, NULL, NULL, 0);

    if (!WinHttpSendRequest(hRequest, WINHTTP_NO_ADDITIONAL_HEADERS, 0, WINHTTP_NO_REQUEST_DATA, 0, 0, 0))
    {
        printf("Error sending request\n");
        return -1;
    }

    if (!WinHttpReceiveResponse(hRequest, NULL)) 
    {
        printf("Error receiving response\n");
        return -1;
    }

    do
    {
        if (!WinHttpQueryDataAvailable(hRequest, &dwSize)) {
            printf("Error getting data available\n");
            return -1;
        }

        pszOutBuffer = malloc(sizeof(char)*dwSize);
        memset(pszOutBuffer, 0, sizeof(char)*dwSize);

        if (!WinHttpReadData(hRequest, (LPVOID)pszOutBuffer, dwSize, &dwDownloaded)) {
            printf("Error reading data\n");
            free(pszOutBuffer);
            return -1;
        }

        printf("%s", pszOutBuffer);

        free(pszOutBuffer);

    } while (dwSize > 0);
}