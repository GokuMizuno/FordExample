#include "fordgreet.h"
#include <stdio.h>

int greet(const char *name, int year, char *out) {
    int n;
    
    n = sprintf(out, "Greetings, %s from %d! I like C, and GO!", name, year);

    return n;
}