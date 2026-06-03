import { useState, useEffect } from 'react'
import { TaskForm } from './components/TaskForm'
import { TaskList } from './components/TaskList'
import type { Task } from './types'

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
      <TaskForm
        categoryId={categoryId}
        title={title}
        notes={notes}
        onCategoryChange={setCategoryId}
        onTitleChange={setTitle}
        onNotesChange={setNotes}
        onSubmit={handleCreate}
      />
      <TaskList tasks={tasks} onComplete={handleComplete} />
    </div>
  )
}

export default App