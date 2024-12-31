import { useQuery } from "@connectrpc/connect-query"
import { listMedia } from "api/media/v1/media-MediaService_connectquery"

export function useMediaList() {
  const { data: mediaList } = useQuery(listMedia)

  return { mediaList }
}
