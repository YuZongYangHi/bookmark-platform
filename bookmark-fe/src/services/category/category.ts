import {request} from 'umi'

export async function listCategory() {
  return request(`/api/v1/category/`)
}

export async function getCategory(id: number) {
  return request(`/api/v1/category/${id}`)
}

export async function updateCategory(id: number, data: object) {
  return request(`/api/v1/category/${id}/`, {
    method: 'PUT',
    data: data,
  })
}

export async function deleteCategory(id: number) {
  return request(`/api/v1/category/${id}/`, {
    method: 'DELETE',
  })
}

export async function createCategory(data: any) {
  return request(`/api/v1/category/`, {
    method: 'POST',
    data: data,
  })
}

