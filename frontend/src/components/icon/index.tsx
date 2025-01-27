import { useState, useEffect, ComponentType } from 'react'

interface IconProps {
  name: string
  className?: string
}

type SvgComponentType = ComponentType<React.SVGProps<SVGSVGElement>>

const userDynamicSvgImport = (name: string) => {
  const [SvgIcon, setSvgIcon] = useState<SvgComponentType | null>(null)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    const importSvgIcon = async () => {
      try {
        const importedIcon = await import(`@/assets/svg/${name}.svg?react`)
        setSvgIcon(() => importedIcon.default)
      } catch (err) {
        setError(err as Error)
      }
    }

    importSvgIcon()
  }, [name])

  return { SvgIcon, error }
}

const Icon: React.FC<IconProps> = ({ name, className }) => {
  const { SvgIcon, error } = userDynamicSvgImport(name)
  if (error) {
    return <div>Error loading icon: {error.message}</div>
  }

  if (!SvgIcon) {
    return <div>Loading...</div>
  }
  return <SvgIcon className={className} />
}

export default Icon
