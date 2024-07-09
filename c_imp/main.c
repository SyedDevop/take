#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void slice(const char *str, char *result, size_t start, size_t end) {
  size_t j = 0;
  for (size_t i = start; i <= end; ++i) {
    result[j++] = str[i];
  }
  result[j++] = 0;
}
void clean(char *p, int *sIdx, int *eIdx) {
  if (eIdx == NULL) {
    printf("eIdx is NULL\n");
    exit(1);
  };

  while (*eIdx > 0 && p[*eIdx - 1] == '/') {
    (*eIdx)--;
  }

  while (p[*sIdx] == '/' || p[*sIdx] == '.') {
    (*sIdx)++;
  }
}

void Dir(char *p, char *path) {
  int sIdx;
  int eIdx = strlen(p);
  if (strlen(p) == 0) {
    path = "";
    return;
  }
  clean(p, &sIdx, &eIdx);
  bool tail_is_file = false;
  for (int i = eIdx - 1; i >= sIdx; i--) {
    if (p[i] == '.') {
      tail_is_file = true;
    }
    if (tail_is_file && p[i] == '/') {
      slice(p, path, sIdx, i - 1);
      return;
    }
  }
  if (tail_is_file) {
    path = "";
    return;
  }

  slice(p, path, sIdx, eIdx);
  return;
}
void Base(char *p, char *file_name) {
  int sIdx;
  int eIdx = strlen(p);
  if (strlen(p) == 0) {
    file_name = "";
    return;
  }
  clean(p, &sIdx, &eIdx);
  bool tail_is_file = false;
  for (int i = eIdx - 1; i >= sIdx; i--) {
    if (p[i] == '.') {
      tail_is_file = true;
    }
    if (p[i] == '/') {
      if (tail_is_file) {
        slice(p, file_name, i + 1, eIdx - 1);
        return;
      }
      file_name = "";
      return;
    }
  }
  if (tail_is_file) {
    slice(p, file_name, sIdx, eIdx - 1);
    return;
  }

  file_name = "";
  return;
}

int main() {
  char *str = "//some/path/with/main.go/";
  char file_name[strlen(str) + 1];
  char path[strlen(str) + 1];
  Base(str, file_name);
  Dir(str, path);
  printf("The Full Path is %s \n", str);
  printf("The file name is %s \n", file_name);
  printf("The Path is %s \n", path);
  return 0;
}
