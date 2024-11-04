#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define USAGE                                                                  \
  "Usage: take [options] <path>\nOptions:\n    -f        Extract and display " \
  "the file name from the given path (e.g., "                                  \
  "fod/bar/bass.go -> bass.go)\n    -d        Extract and display the "        \
  "directory path from the given path "                                        \
  "(e.g., fod/bar/bass.go -> fod/bar)\n    -h        Show help and usage "     \
  "information\nExample:\n   take -f fod/bar/bass.go\n   take -d "             \
  "fod/bar/bass.go\n"

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

int main(int argc, char *argv[]) {
  if (argc == 2) {
    if (strcmp(argv[1], "-h") == 0) {
      printf("%s ", USAGE);
      return 0;
    }
  }
  if (argc == 3) {
    char *str = argv[2];
    if (strcmp(argv[1], "-h") == 0) {
      printf("%s ", USAGE);
      return 0;
    }
    if (strcmp(argv[1], "-f") == 0) {
      char file_name[strlen(str) + 1];
      Base(str, file_name);
      printf("%s", file_name);
      return 0;
    }
    if (strcmp(argv[1], "-d") == 0) {
      char path[strlen(str) + 1];
      Dir(str, path);
      printf("%s", path);
      return 0;
    }
  }
  return 0;
}
