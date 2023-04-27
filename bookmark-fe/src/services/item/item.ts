import {request} from 'umi'

export async function listItem() {
  return request(`/api/v1/item/`)
}

export async function getItem(id: number) {
  return request(`/api/v1/item/${id}`)
}

export async function updateItem(id: number, data: object) {
  return request(`/api/v1/item/${id}/`, {
    method: 'PUT',
    data: data,
  })
}

export async function deleteItem(id: number) {
  return request(`/api/v1/item/${id}/`, {
    method: 'DELETE',
  })
}

export async function createItem(data: any) {
  return request(`/api/v1/item/`, {
    method: 'POST',
    data: data,
  })
}

