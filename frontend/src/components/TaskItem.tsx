import type { Task } from '../types'

type Props = {
  task: Task
  onComplete: (id: number) => void
}

export function TaskItem({ task, onComplete }: Props) {
  return (
    <li>
      [{task.category?.name}] {task.title} - {task.completed ? '完了' : (
        <button onClick={() => onComplete(task.id)}>完了にする</button>
      )}
    </li>
  )
}