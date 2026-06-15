import React from 'react';
import TaskCard from './TaskCard';

export default function TaskList({ tasks, groups, onToggle, onDelete, onReminder }) {
  const groupById = Object.fromEntries(groups.map((g) => [g.id, g]));
  if (!tasks.length) return <div className="empty">No tasks yet.</div>;
  return <section className="task-list">{tasks.map((task)=><TaskCard key={task.id} task={task} group={groupById[task.group_id]} onToggle={onToggle} onDelete={onDelete} onReminder={onReminder} />)}</section>;
}
