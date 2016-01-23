#include "cfolt.h"
#include <stdio.h>
int main () {

  GoString str = {".",0};
  int i = LoadGeneratorFromPath(str);
  printf("Got int %i", i);
  return 0;
}