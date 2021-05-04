#include <stdio.h>

#include <iostream>

#include <windows.h>
#include <winhttp.h>

#include "lp.hpp"

using namespace std;



int main(int argc, char **argv)
{
    ListeningPost lp("ipinfo.io", 80, "curl/7.64.1");

    try
    {
        cout << lp.getData() << endl;
    } catch (char const *e) {
        cout << "Error: " << e << endl;
    }

}