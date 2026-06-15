import React, { useState } from 'react';

export default function TaskCard({ task, group, onToggle, onDelete, onReminder }) {
  const [remindAt, setRemindAt] = useState('');
  const done = task.status === 'completed';
  const submitReminder = (e) => { e.preventDefault(); if (!remindAt) return; onReminder({ task_id: task.id, remind_at: new Date(remindAt).toISOString(), message: task.title }); setRemindAt(''); };
  return <article className={done ? 'task-card done' : 'task-card'}><div className="task-main"><input type="checkbox" checked={done} onChange={()=>onToggle(task)} /><div><h3>{task.title}</h3>{task.description && <p>{task.description}</p>}<div className="meta"><span className={`priority ${task.priority}`}>{task.priority}</span>{group && <span>{group.name}</span>}{task.due_at && <span>{new Date(task.due_at).toLocaleString()}</span>}</div></div></div><form className="reminder-row" onSubmit={submitReminder}><input type="datetime-local" value={remindAt} onChange={(e)=>setRemindAt(e.target.value)} /><button type="submit">Remind</button><button type="button" className="danger" onClick={()=>onDelete(task.id)}>Delete</button></form></article>;
}
