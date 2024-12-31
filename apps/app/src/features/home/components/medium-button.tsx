import { useRouter } from "expo-router"
import { Button, Pressable, Text } from "react-native"
import { StyleSheet } from "react-native-unistyles"

interface MediumButtonProps {
  mediumId: string
  title: string
}

export function MediumButton(props: MediumButtonProps) {
  const router = useRouter()

  return (
    <Pressable
      onPress={() => router.navigate(`/media/${props.mediumId}`)}
      style={styles.button}
    >
      <Text>{props.title}</Text>
    </Pressable>
  )
}

const styles = StyleSheet.create({
  button: {
    backgroundColor: "red",
  },
})
