import { useQuery } from "@connectrpc/connect-query"
import { getMedium } from "api/media/v1/media-MediaService_connectquery"

export function useMedium(mediumId: string) {
  const { data: medium } = useQuery(getMedium, { id: mediumId })

  return { medium }
}
