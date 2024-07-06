export type UserInQueue = {
  created_at: string
  twitch_user_id: string
  twitch_username: string
  notes: string
}

export type InfoResponse = {
  total: number
  users: UserInQueue[]
}

export const getUsersInQueue = async (): Promise<InfoResponse> => {
  const res = await fetch('/api/info')
  const resJson = await res.json() as InfoResponse

  return resJson
}
