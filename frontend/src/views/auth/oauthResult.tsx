/*
 * @fileName oauthResult.tsx
 * @author Di Sheng
 * @date 2024/08/19 11:10:13
 * @description OAuth result page
 */

import { useEffect } from 'react'

const OAuthResult = () => {
  useEffect(() => {
    // const params = new URLSearchParams(window.location.search)
    // const token = params.get('token')
    // const result = params.get('result')

    // if (window.opener && result === 'success') {
    //   window.opener.postMessage(
    //     { token, status: 'success' },
    //     window.location.origin
    //   )
    // }

    if (window.opener) {
      window.opener.postMessage({ status: 'success' }, window.location.origin)
    }

    window.close()
  }, [])
  return <div>Login successful. You can close this window.</div>
}

export default OAuthResult
