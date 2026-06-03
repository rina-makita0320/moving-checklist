import type { Task } from '../types'
import { TaskItem } from './TaskItem'

type Props = {
  tasks: Task[]
  onComplete: (id: number) => void
}

export function TaskList({ tasks, onComplete }: Props) {
  return (
    <ul>
      {tasks.map(task => (
        <TaskItem key={task.id} task={task} onComplete={onComplete} />
      ))}
    </ul>
  )
}