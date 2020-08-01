#define _GNU_SOURCE

#include <dlfcn.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

void ProgressWrite(size_t);

static ssize_t (*original_write)(int, const void*, size_t n);

void progress_exit(int status, void* arg) {
  ProgressWrite(0);
}

void init_write() {
  original_write = dlsym(RTLD_NEXT, "write");
  on_exit(progress_exit, NULL);
}

ssize_t write (int fd, const void *buf, size_t n) {
  if (fd == STDOUT_FILENO)
    ProgressWrite(n);
  return original_write(fd, buf, n);
}

