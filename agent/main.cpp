#include <stdio.h>

#include <iostream>

#include <windows.h>
#include <winhttp.h>

#include "lp.hpp"

using namespace std;



int main(int argc, char **argv)
{
    if (argc < 3) {
        cout << "Usage: " << argv[0] << " HOSTNAME PORT" << endl;
        return -1;
    }

    ListeningPost lp(argv[1], atoi(argv[2]), "curl/7.64.1");

    try
    {
        cout << lp.getData() << endl;
    } catch (runtime_error e) {
        cout << "Error: " << e.what() << endl;
    }

    return 0;

}