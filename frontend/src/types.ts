export type Category = {
  id: number
  name: string
}

export type Task = {
  id: number
  category_id: number
  category: Category
  title: string
  due_date: string | null
  completed: boolean
  notes: string
  created_at: string
}