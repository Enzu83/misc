import sys
import math
import matplotlib.pyplot as plt

def is_prime(n):
    if n <= 1:
        return 1
    
    for k in range(2, math.floor(math.sqrt(n)) + 1):
        if n % k == 0:
            return k
        
    return None

def prime_dist(n):
    primes = set()

    for k in range(3, n+1):
        if is_prime(k) is None:
            primes.add(k)
    
    return primes

def print_primes(primes):
    sup = max(primes)+1
    x = [k for k in range(sup)]
    y = [0]

    for k in range(1, sup):
        if k in primes:
            y.append(k - len(y)-1)
        else:
            y.append(y[-1] + 1)

    plt.plot(x, y)
    plt.show()

def main():
    primes = (prime_dist(int(sys.argv[1])))

    print_primes(primes)

if __name__ == "__main__":
    main()