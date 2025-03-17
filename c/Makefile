CC = gcc
CFLAGS = -Wall -Wextra -g -O2 -fopenmp
OBJ = main.o iterative.o recursive.o
TARGET = main

all: $(TARGET)

$(TARGET): $(OBJ)
	$(CC) $(CFLAGS) -o $@ $^

main.o: main.c iterative.h recursive.h
	$(CC) $(CFLAGS) -c -o $@ $<

iterative.o: iterative.c iterative.h
	$(CC) $(CFLAGS) -c -o $@ $<

recursive.o: recursive.c recursive.h
	$(CC) $(CFLAGS) -c -o $@ $<

clean:
	rm -f $(OBJ) $(TARGET) results

run: $(TARGET)
	./$(TARGET)