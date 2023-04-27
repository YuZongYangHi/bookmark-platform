import {request} from 'umi'

export async function lisMenu() {
  return request(`/api/v1/menu/list`)
}
