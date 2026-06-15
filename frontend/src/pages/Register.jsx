import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function Register() {
  const { register } = useAuth();
  const navigate = useNavigate();
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const [error, setError] = useState('');
  const submit = async (e) => { e.preventDefault(); setError(''); try { await register(form); navigate('/'); } catch (err) { setError(err.response?.data?.error || 'Registration failed'); } };
  return <main className="auth-screen"><form className="auth-panel" onSubmit={submit}><h1>Create account</h1><p>Start organizing your work in a few seconds.</p>{error && <div className="alert">{error}</div>}<label>Name<input value={form.name} onChange={(e)=>setForm({...form,name:e.target.value})} required /></label><label>Email<input type="email" value={form.email} onChange={(e)=>setForm({...form,email:e.target.value})} required /></label><label>Password<input type="password" minLength="6" value={form.password} onChange={(e)=>setForm({...form,password:e.target.value})} required /></label><button type="submit">Create account</button><span>Already registered? <Link to="/login">Sign in</Link></span></form></main>;
}
