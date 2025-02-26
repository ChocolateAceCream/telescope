export function decodeBase64(
  base64Str: string | undefined
): Record<string, any> {
  try {
    if (!base64Str) {
      return {} // Return empty object if input is empty
    }

    // Decode Base64
    const decodedStr =
      typeof atob !== 'undefined'
        ? atob(base64Str) // Browser environment
        : Buffer.from(base64Str, 'base64').toString('utf-8') // Node.js

    // Parse JSON
    return JSON.parse(decodedStr)
  } catch (error) {
    console.error('Error decoding Base64:', error)
    return {} // Return empty object if decoding or parsing fails
  }
}
