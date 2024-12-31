import { useLocalSearchParams } from "expo-router"
import { MediumPage } from "../../features/media/medium-page"

export default function () {
  const { mediumId } = useLocalSearchParams()

  if (typeof mediumId !== "string") {
    throw new Error("Expected mediumId to be a string")
  }

  return <MediumPage mediumId={mediumId} />
}
