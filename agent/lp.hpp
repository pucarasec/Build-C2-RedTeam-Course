
#ifndef __LISTENING_POST_HPP
#define __LISTENING_POST_HPP

#include <string>

using namespace std;

class ListeningPost
{
    public:
    ListeningPost(const char *hostname, int port, const char *userAgent);
    string getData();

    private:
    wstring hostname;
    int port;
    wstring userAgent;
};

#endif