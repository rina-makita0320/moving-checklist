type Props = {
  categoryId: number
  title: string
  notes: string
  onCategoryChange: (id: number) => void
  onTitleChange: (title: string) => void
  onNotesChange: (notes: string) => void
  onSubmit: () => void
}

export function TaskForm({
  categoryId,
  title,
  notes,
  onCategoryChange,
  onTitleChange,
  onNotesChange,
  onSubmit,
}: Props) {
  return (
    <div>
      <select value={categoryId} onChange={e => onCategoryChange(Number(e.target.value))}>
        <option value={1}>行政系</option>
        <option value={2}>ライフライン系</option>
      </select>
      <input
        type="text"
        placeholder="手続き名"
        value={title}
        onChange={e => onTitleChange(e.target.value)}
      />
      <input
        type="text"
        placeholder="メモ"
        value={notes}
        onChange={e => onNotesChange(e.target.value)}
      />
      <button onClick={onSubmit}>追加</button>
    </div>
  )
}