import { TransportProvider } from "@connectrpc/connect-query"
import { createConnectTransport } from "@connectrpc/connect-web"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { Stack } from "expo-router"
import { SafeAreaView } from "react-native"

const finalTransport = createConnectTransport({
  baseUrl: "http://localhost:8080",
})

const queryClient = new QueryClient()

export default function RootLayout() {
  return (
    <SafeAreaView style={{ flex: 1 }}>
      <TransportProvider transport={finalTransport}>
        <QueryClientProvider client={queryClient}>
          <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen name="(authenticated)" />
            <Stack.Screen name="media/[mediumId]" />
          </Stack>
        </QueryClientProvider>
      </TransportProvider>
    </SafeAreaView>
  )
}
