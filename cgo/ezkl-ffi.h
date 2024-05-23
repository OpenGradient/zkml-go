/* ezkl-ffi Header */
#ifdef __cplusplus
extern "C" {{
#endif


#ifndef EZKL_FFI_H
#define EZKL_FFI_H

/* Generated with cbindgen:0.26.0 */

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

bool verify_proof(size_t proof_length,
                  const uint8_t *proof,
                  size_t vk_length,
                  const uint8_t *vk,
                  size_t settings_length,
                  const uint8_t *settings,
                  size_t srs_length,
                  const uint8_t *srs);

const char *prove(size_t witness_length,
                  const uint8_t *witness,
                  size_t pk_length,
                  const uint8_t *pk,
                  size_t cpmpiled_circuit_length,
                  const uint8_t *compiled_circuit,
                  size_t srs_length,
                  const uint8_t *srs);

void gen_vk(size_t compiled_circuit_length,
            const uint8_t *compiled_circuit,
            size_t params_serialized_length,
            const uint8_t *params_serialized,
            bool compress_selectors);

void gen_pk(size_t vk_length,
            const uint8_t *vk,
            size_t circuit_length,
            const uint8_t *circuit,
            size_t params_serialized_length,
            const uint8_t *params_serialized);

#endif /* EZKL_FFI_H */

#ifdef __cplusplus
} /* extern "C" */
#endif
