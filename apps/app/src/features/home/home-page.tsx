import { View } from "react-native"
import { useMediaList } from "../media/hooks/use-media-list"
import { MediumButton } from "./components/medium-button"

export function HomePage() {
  const { mediaList } = useMediaList()

  return (
    <View>
      {mediaList?.media.map((media) => (
        <MediumButton key={media.id} mediumId={media.id} title={media.title} />
      ))}
    </View>
  )
}
