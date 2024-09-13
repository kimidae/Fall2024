def sieve_of_eratosthenes(limit):
    is_prime = [True] * (limit + 1)
    p = 2
    while (p * p <= limit):
        if is_prime[p]:
            for i in range(p * p, limit + 1, p):
                is_prime[i] = False
        p += 1
    primes = [p for p in range(2, limit + 1) if is_prime[p]]
    return primes

def find_superprimes(n):
    # Estimate the upper limit to find enough primes
    limit = 100000  # We might adjust this if necessary
    primes = sieve_of_eratosthenes(limit)
    
    # Find superprimes
    superprimes = []
    for i in range(1, len(primes) + 1):
        if i < len(primes):
            prime_index = i
            if prime_index > 1 and primes[prime_index - 1] in primes:  # Check if index is a prime
                superprimes.append(primes[prime_index - 1])
    
    return superprimes[n - 1]

def main():
    import sys
    input = sys.stdin.read
    n = int(input().strip())
    
    result = find_superprimes(n)
    print(result)

if __name__ == "__main__":
    main()
