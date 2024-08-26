#include <immintrin.h>

unsigned long long SlowMultiplication(unsigned long long first, unsigned long long second)
{
    unsigned long long iteration;
    unsigned long long quantity;
    unsigned long long remainder;

    const unsigned long long parallelization = 64;

    __m256i added;
    __m256i product;

    added[0] = first;
    added[1] = first;
    added[2] = first;
    added[3] = first;

    product[0] = 0;
    product[1] = 0;
    product[2] = 0;
    product[3] = 0;

    quantity = second / parallelization;
    remainder = second % parallelization;

    for (iteration = 0; iteration < remainder; iteration++)
    {
        product[0] += first;
    }

    for (iteration = 0; iteration < quantity; iteration++)
    {
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
        product = _mm256_add_epi64(product, added);
    }

    return product[0] + product[1] + product[2] + product[3];
}

// Only for numbers without fractional parts.
double SlowMultiplicationF64(double first, double second)
{
    unsigned long long iteration;
    unsigned long long quantity;
    unsigned long long remainder;

    // Increasing parallelism does not lead to increased performance
    const unsigned long long parallelization = 4;

    __m256d added;
    __m256d product;

    added[0] = first;
    added[1] = first;
    added[2] = first;
    added[3] = first;

    product[0] = 0;
    product[1] = 0;
    product[2] = 0;
    product[3] = 0;

    quantity = (unsigned long long)(second) / parallelization;
    remainder = (unsigned long long)(second) % parallelization;

    for (iteration = 0; iteration < remainder; iteration++)
    {
        product[0] += first;
    }

    for (iteration = 0; iteration < quantity; iteration++)
    {
        product = _mm256_add_pd(product, added);
    }

    return product[0] + product[1] + product[2] + product[3];
}