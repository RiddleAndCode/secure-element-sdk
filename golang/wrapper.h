 #include <stdio.h>
#include <atca_status.h>
#include <atca_device.h>
#include <atca_command.h>
#include <atca_iface.h>
#include <atca_cfgs.h>
#include <atca_host.h>
#include <atca_execution.h>
#include <atca_basic.h>
#include <atca_helpers.h>

#define ATCAPRINTF

#define PUBLIC_KEY_SIZE 64
#define PRIVATE_KEY_SIZE 32
#define SIGNATURE_SIZE 64

#define __SUCCESS__


ATCAIfaceCfg g_iface_config = {
    .iface_type        = ATCA_I2C_IFACE,
    .devtype           = ATECC508A,
    .atcai2c           = {
        .slave_address = 0xC0,
        .bus           = 1,
        .baud          = 400000,
    },
    .wake_delay        = 1500,
    .rx_retries        = 20
};

ATCAIfaceCfg *gCfg = &g_iface_config;


/** \brief Executes Random command, which generates a 32 byte random number
 *          from the CryptoAuth device.
 *
 * \param[out] randombytes  32 bytes of random data is returned here.
 *
 * \return __SUCCESS__ on success, otherwise an error code.
 */
int getRandomNumber( uint8_t* randombytes ){
    ATCA_STATUS status;
    status =  atcab_init(gCfg);
    status =  atcab_random(randombytes);
    status =  atcab_release();
    return (int)status;
}


/** \brief Uses GenKey command to calculate the public key from an existing
 *          private key in a slot.
 *
 *  \param[out] public_key  Public key will be returned here. Format will be
 *                          the X and Y integers in big-endian format.
 *                          64 bytes for P256 curve. Set to NULL if public key
 *                          isn't required.
 *
 *  \return __SUCCESS__ on success, otherwise an error code.
 */
/*
     uint8_t public_key_bob[ATCA_PUB_KEY_SIZE];
*/
int getPublicKey( uint8_t* public_key ){
    ATCA_STATUS status;
    status =  atcab_init(gCfg);
    status =  atcab_get_pubkey(0, public_key);
    status =  atcab_release();
    return (int)status;
}

/** \brief Executes Sign command, to sign a 32-byte external message using the
 *                   private key in the specified slot. The message to be signed
 *                   will be loaded into the Message Digest Buffer to the
 *                   ATECC608A device or TempKey for other devices.
 *
 *  \param[in]  digest        32-byte message to be signed. Typically the SHA256
 *                         hash of the full message.
 *  \param[out] signature  Signature will be returned here. Format is R and S
 *                         integers in big-endian format. 64 bytes for P256
 *                         curve.
 *
 * \return __SUCCESS__ on success, otherwise an error code.
 */
int sign( uint8_t* digest, uint8_t* signature ){
    ATCA_STATUS status;
    status =  atcab_init(gCfg);
    status =  atcab_sign(0, digest, signature);
    status =  atcab_release();
    return (int)status;
}


/** \brief Executes the Verify command, which verifies a signature (ECDSA
 *          verify operation) with all components (message, signature, and
 *          public key) supplied. The message to be signed will be loaded into
 *          the Message Digest Buffer to the ATECC608A device or TempKey for
 *          other devices.
 *
 * \param[in]  digest      32 byte message to be verified. Typically
 *                          the SHA256 hash of the full message.
 * \param[in]  signature    Signature to be verified. R and S integers in
 *                          big-endian format. 64 bytes for P256 curve.
 * \param[in]  public_key   The public key to be used for verification. X and
 *                          Y integers in big-endian format. 64 bytes for
 *                          P256 curve.
 * \param[out] is_verified  Boolean whether or not the message, signature,
 *                          public key verified.
 *
 * \return __SUCCESS__ on verification success or failure, because the
 *         command still completed successfully.
 */
int verify_extern( uint8_t* digest, uint8_t* signature, uint8_t* public_key, bool* is_verified){
    ATCA_STATUS status;
    status =  atcab_init(gCfg);
    status =  atcab_verify_extern(digest, signature, public_key, is_verified);
    status =  atcab_release();
    return (int) status;
}

