#ifndef _fordgreet_h
#define _fordgreet_h

struct Greetee {
    const char *name;
    int year;
};

int greet(struct Greetee *g, char *out);

#endif