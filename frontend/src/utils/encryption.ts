/*
 * @fileName encryption.ts
 * @author Di Sheng
 * @date 2024/09/09 20:18:06
 * @description encryption util functions
 */

// src/utils/hashUtils.ts
import SHA256 from 'crypto-js/sha256'
import encUtf8 from 'crypto-js/enc-utf8'

export function sha256(message: string): string {
  const r = SHA256(encUtf8.parse(message)).toString()
  console.log('SHA256:', r)
  return r
}
