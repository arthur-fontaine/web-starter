import { Text, View } from "react-native"
import { useMedium } from "./hooks/use-medium"

interface MediaPageProps {
  mediumId: string
}

export function MediumPage(props: MediaPageProps) {
  const { medium } = useMedium(props.mediumId)

  return (
    <View>
      <Text>{medium?.title}</Text>
    </View>
  )
}
