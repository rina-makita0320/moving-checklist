import { useState, useEffect } from 'react'

type Category = {
  id: number
  name: string
}

type Task = {
  id: number
  category_id: number
  category: Category
  title: string
  due_date: string | null
  completed: boolean
  notes: string
  created_at: string
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [title, setTitle] = useState('')
  const [notes, setNotes] = useState('')
  const [categoryId, setCategoryId] = useState(1)

  const fetchTasks = () => {
    fetch('http://localhost:8080/tasks')
      .then(res => res.json())
      .then(data => setTasks(data))
  }

  useEffect(() => {
    fetchTasks()
  }, [])

  const handleCreate = () => {
    fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ category_id: categoryId, title, notes })
    }).then(() => {
      setTitle('')
      setNotes('')
      fetchTasks()
    })
  }

  const handleComplete = (id: number) => {
    fetch(`http://localhost:8080/tasks/${id}/complete`, {
      method: 'PATCH'
    }).then(() => fetchTasks())
  }

  return (
    <div>
      <h1>引越しチェックリスト</h1>

      <div>
        <select value={categoryId} onChange={e => setCategoryId(Number(e.target.value))}>
          <option value={1}>行政系</option>
          <option value={2}>ライフライン系</option>
        </select>
        <input
          type="text"
          placeholder="手続き名"
          value={title}
          onChange={e => setTitle(e.target.value)}
        />
        <input
          type="text"
          placeholder="メモ"
          value={notes}
          onChange={e => setNotes(e.target.value)}
        />
        <button onClick={handleCreate}>追加</button>
      </div>

      <ul>
        {tasks.map(task => (
          <li key={task.id}>
            [{task.category?.name}] {task.title} - {task.completed ? '完了' : (
              <button onClick={() => handleComplete(task.id)}>完了にする</button>
            )}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App