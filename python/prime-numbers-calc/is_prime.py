import math
import sys

def is_prime(n):
    if n <= 1:
        return 1
    
    for k in range(2, 1 + math.floor(math.sqrt(n)) + 1):
        if n % k == 0:
            return k
        
    return None

def main():
    number = 2 ** int(sys.argv[1]) - 1
    
    divisor = is_prime(number)

    if divisor is None:
        print(f"{number} is prime")
    else:
        print(f"{number} is divisible by {divisor}")

if __name__ == "__main__":
    main()