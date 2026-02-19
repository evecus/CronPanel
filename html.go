package main

var indexHTML = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>CronPanel — 定时任务管理</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=DM+Mono:wght@400;500&family=Syne:wght@400;600;700;800&family=Inter:wght@300;400;500&display=swap" rel="stylesheet">
<style>
  :root {
    --bg: #0a0b0f;
    --surface: #111318;
    --surface2: #181c24;
    --border: #1f2433;
    --border2: #2a3045;
    --accent: #6c63ff;
    --accent2: #a855f7;
    --accent3: #06b6d4;
    --green: #10b981;
    --red: #f43f5e;
    --orange: #f97316;
    --text: #e2e8f0;
    --text2: #94a3b8;
    --text3: #475569;
    --glow: rgba(108,99,255,0.15);
  }
  * { box-sizing: border-box; margin: 0; padding: 0; }
  
  body {
    background: var(--bg);
    color: var(--text);
    font-family: 'Inter', sans-serif;
    font-size: 14px;
    min-height: 100vh;
    overflow-x: hidden;
  }

  /* Animated grid background */
  body::before {
    content: '';
    position: fixed;
    inset: 0;
    background-image: 
      linear-gradient(rgba(108,99,255,0.03) 1px, transparent 1px),
      linear-gradient(90deg, rgba(108,99,255,0.03) 1px, transparent 1px);
    background-size: 40px 40px;
    pointer-events: none;
    z-index: 0;
  }

  .layout {
    display: flex;
    min-height: 100vh;
    position: relative;
    z-index: 1;
  }

  /* SIDEBAR */
  .sidebar {
    width: 220px;
    background: var(--surface);
    border-right: 1px solid var(--border);
    display: flex;
    flex-direction: column;
    padding: 24px 0;
    position: fixed;
    top: 0; left: 0; bottom: 0;
    z-index: 100;
  }

  .logo {
    padding: 0 20px 28px;
    border-bottom: 1px solid var(--border);
  }
  .logo-icon {
    width: 36px; height: 36px;
    background: linear-gradient(135deg, var(--accent), var(--accent2));
    border-radius: 10px;
    display: flex; align-items: center; justify-content: center;
    font-size: 18px;
    margin-bottom: 10px;
    box-shadow: 0 0 20px var(--glow);
  }
  .logo-name {
    font-family: 'Syne', sans-serif;
    font-weight: 800;
    font-size: 16px;
    color: var(--text);
    letter-spacing: 0.5px;
  }
  .logo-sub { font-size: 11px; color: var(--text3); margin-top: 2px; }

  .nav { padding: 20px 12px; flex: 1; }
  .nav-label {
    font-size: 10px; font-weight: 600;
    text-transform: uppercase; letter-spacing: 1.5px;
    color: var(--text3);
    padding: 0 8px;
    margin-bottom: 8px;
  }
  .nav-item {
    display: flex; align-items: center; gap: 10px;
    padding: 10px 12px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 13px;
    color: var(--text2);
    transition: all 0.15s;
    margin-bottom: 2px;
  }
  .nav-item:hover { background: var(--surface2); color: var(--text); }
  .nav-item.active {
    background: linear-gradient(135deg, rgba(108,99,255,0.2), rgba(168,85,247,0.1));
    color: #a78bfa;
    border: 1px solid rgba(108,99,255,0.3);
  }
  .nav-icon { font-size: 16px; width: 20px; text-align: center; }

  .sidebar-footer {
    padding: 16px 20px;
    border-top: 1px solid var(--border);
    font-size: 11px;
    color: var(--text3);
  }
  .status-dot {
    display: inline-block;
    width: 6px; height: 6px;
    background: var(--green);
    border-radius: 50%;
    margin-right: 6px;
    box-shadow: 0 0 6px var(--green);
    animation: pulse 2s infinite;
  }
  @keyframes pulse { 0%,100%{opacity:1} 50%{opacity:0.5} }

  /* MAIN */
  .main {
    margin-left: 220px;
    flex: 1;
    padding: 32px 36px;
    max-width: 1200px;
  }

  .page-header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    margin-bottom: 32px;
    gap: 16px;
  }
  .page-title {
    font-family: 'Syne', sans-serif;
    font-size: 28px;
    font-weight: 800;
    background: linear-gradient(135deg, var(--text) 40%, var(--accent2));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }
  .page-sub { font-size: 13px; color: var(--text3); margin-top: 4px; }

  .btn {
    display: inline-flex; align-items: center; gap: 8px;
    padding: 10px 18px;
    border-radius: 8px;
    border: none;
    cursor: pointer;
    font-size: 13px;
    font-weight: 500;
    font-family: 'Inter', sans-serif;
    transition: all 0.2s;
    text-decoration: none;
  }
  .btn-primary {
    background: linear-gradient(135deg, var(--accent), var(--accent2));
    color: white;
    box-shadow: 0 4px 20px rgba(108,99,255,0.3);
  }
  .btn-primary:hover {
    transform: translateY(-1px);
    box-shadow: 0 6px 28px rgba(108,99,255,0.4);
  }
  .btn-ghost {
    background: var(--surface2);
    color: var(--text2);
    border: 1px solid var(--border2);
  }
  .btn-ghost:hover { background: var(--border2); color: var(--text); }
  .btn-danger { background: rgba(244,63,94,0.15); color: var(--red); border: 1px solid rgba(244,63,94,0.3); }
  .btn-danger:hover { background: rgba(244,63,94,0.25); }
  .btn-sm { padding: 6px 12px; font-size: 12px; }

  /* STATS */
  .stats-row {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 16px;
    margin-bottom: 28px;
  }
  .stat-card {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 12px;
    padding: 20px;
    position: relative;
    overflow: hidden;
  }
  .stat-card::before {
    content: '';
    position: absolute;
    top: 0; right: 0;
    width: 80px; height: 80px;
    border-radius: 50%;
    transform: translate(25%, -25%);
    opacity: 0.15;
  }
  .stat-card.total::before { background: var(--accent); }
  .stat-card.active::before { background: var(--green); }
  .stat-card.disabled::before { background: var(--text3); }
  .stat-label { font-size: 11px; color: var(--text3); text-transform: uppercase; letter-spacing: 1px; }
  .stat-value {
    font-family: 'Syne', sans-serif;
    font-size: 32px; font-weight: 800;
    margin: 8px 0 4px;
    color: var(--text);
  }
  .stat-sub { font-size: 12px; color: var(--text2); }

  /* JOB LIST */
  .section-header {
    display: flex; align-items: center; justify-content: space-between;
    margin-bottom: 16px;
  }
  .section-title {
    font-family: 'Syne', sans-serif;
    font-size: 16px; font-weight: 700;
  }

  .job-list { display: flex; flex-direction: column; gap: 10px; }

  .job-card {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 12px;
    padding: 16px 20px;
    display: flex;
    align-items: center;
    gap: 16px;
    transition: all 0.2s;
    animation: slideIn 0.3s ease;
  }
  @keyframes slideIn { from { opacity:0; transform:translateY(8px); } to { opacity:1; transform:none; } }
  .job-card:hover { border-color: var(--border2); background: var(--surface2); }
  .job-card.disabled { opacity: 0.5; }

  .job-status {
    width: 8px; height: 8px;
    border-radius: 50%;
    flex-shrink: 0;
  }
  .job-status.on { background: var(--green); box-shadow: 0 0 8px var(--green); }
  .job-status.off { background: var(--text3); }

  .job-info { flex: 1; min-width: 0; }
  .job-comment {
    font-size: 13px; font-weight: 500;
    color: var(--text);
    margin-bottom: 6px;
    white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
  }
  .job-command {
    font-family: 'DM Mono', monospace;
    font-size: 11px;
    color: var(--accent3);
    background: rgba(6,182,212,0.08);
    border: 1px solid rgba(6,182,212,0.15);
    border-radius: 5px;
    padding: 3px 8px;
    display: inline-block;
    max-width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .job-schedule {
    font-family: 'DM Mono', monospace;
    font-size: 12px;
    color: var(--accent2);
    background: rgba(168,85,247,0.1);
    border: 1px solid rgba(168,85,247,0.2);
    padding: 4px 10px;
    border-radius: 20px;
    white-space: nowrap;
    flex-shrink: 0;
  }

  .job-actions { display: flex; gap: 8px; flex-shrink: 0; }

  .empty-state {
    text-align: center;
    padding: 60px 20px;
    color: var(--text3);
  }
  .empty-icon { font-size: 48px; margin-bottom: 16px; opacity: 0.5; }
  .empty-text { font-size: 14px; }

  /* MODAL */
  .modal-overlay {
    position: fixed; inset: 0;
    background: rgba(0,0,0,0.7);
    backdrop-filter: blur(4px);
    z-index: 1000;
    display: none;
    align-items: center;
    justify-content: center;
    padding: 20px;
  }
  .modal-overlay.open { display: flex; }

  .modal {
    background: var(--surface);
    border: 1px solid var(--border2);
    border-radius: 16px;
    width: 100%;
    max-width: 600px;
    max-height: 90vh;
    overflow-y: auto;
    box-shadow: 0 25px 60px rgba(0,0,0,0.5), 0 0 0 1px rgba(108,99,255,0.1);
    animation: modalIn 0.25s ease;
  }
  @keyframes modalIn { from { opacity:0; transform:scale(0.95) translateY(10px); } to { opacity:1; transform:none; } }

  .modal-header {
    padding: 24px 24px 20px;
    border-bottom: 1px solid var(--border);
    display: flex; align-items: center; justify-content: space-between;
  }
  .modal-title {
    font-family: 'Syne', sans-serif;
    font-size: 18px; font-weight: 700;
  }
  .modal-close {
    background: none; border: none; cursor: pointer;
    color: var(--text3); font-size: 20px;
    width: 32px; height: 32px;
    border-radius: 6px;
    display: flex; align-items: center; justify-content: center;
    transition: all 0.15s;
  }
  .modal-close:hover { background: var(--surface2); color: var(--text); }

  .modal-body { padding: 24px; }
  .modal-footer {
    padding: 16px 24px;
    border-top: 1px solid var(--border);
    display: flex; gap: 10px; justify-content: flex-end;
  }

  /* FORM */
  .form-group { margin-bottom: 20px; }
  .form-label {
    display: block;
    font-size: 12px; font-weight: 500;
    color: var(--text2);
    margin-bottom: 8px;
    text-transform: uppercase;
    letter-spacing: 0.8px;
  }
  .form-input, .form-select, .form-textarea {
    width: 100%;
    background: var(--surface2);
    border: 1px solid var(--border2);
    border-radius: 8px;
    padding: 10px 14px;
    color: var(--text);
    font-size: 13px;
    font-family: 'Inter', sans-serif;
    outline: none;
    transition: border-color 0.15s;
  }
  .form-input:focus, .form-select:focus, .form-textarea:focus {
    border-color: var(--accent);
    box-shadow: 0 0 0 3px rgba(108,99,255,0.1);
  }
  .form-input::placeholder { color: var(--text3); }
  .form-select { cursor: pointer; }
  .form-select option { background: var(--surface2); }
  .form-textarea {
    min-height: 100px; resize: vertical;
    font-family: 'DM Mono', monospace;
    font-size: 12px;
  }

  /* Mode selector */
  .mode-tabs {
    display: flex;
    background: var(--surface2);
    border: 1px solid var(--border2);
    border-radius: 10px;
    padding: 4px;
    gap: 2px;
    margin-bottom: 20px;
    flex-wrap: wrap;
  }
  .mode-tab {
    flex: 1;
    min-width: 80px;
    padding: 8px 10px;
    border-radius: 6px;
    border: none;
    background: none;
    color: var(--text3);
    font-size: 12px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s;
    font-family: 'Inter', sans-serif;
    text-align: center;
  }
  .mode-tab:hover { color: var(--text2); }
  .mode-tab.active {
    background: var(--accent);
    color: white;
    box-shadow: 0 2px 8px rgba(108,99,255,0.4);
  }

  .time-row { display: flex; gap: 12px; }
  .time-row .form-group { flex: 1; }

  .mode-section { display: none; }
  .mode-section.active { display: block; }

  .cron-preview {
    background: var(--bg);
    border: 1px solid var(--border);
    border-radius: 8px;
    padding: 12px 16px;
    font-family: 'DM Mono', monospace;
    font-size: 13px;
    color: var(--accent2);
    margin-bottom: 20px;
  }
  .cron-preview-label { font-size: 10px; color: var(--text3); margin-bottom: 4px; font-family: 'Inter'; }

  .tag { 
    display: inline-flex; align-items: center;
    padding: 2px 8px; border-radius: 4px;
    font-size: 11px; font-weight: 500;
  }
  .tag-green { background: rgba(16,185,129,0.15); color: var(--green); }
  .tag-red { background: rgba(244,63,94,0.15); color: var(--red); }

  .loading { text-align: center; padding: 40px; color: var(--text3); }
  .spinner {
    width: 24px; height: 24px;
    border: 2px solid var(--border2);
    border-top-color: var(--accent);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto 12px;
  }
  @keyframes spin { to { transform: rotate(360deg); } }

  .toast {
    position: fixed;
    bottom: 24px; right: 24px;
    background: var(--surface);
    border: 1px solid var(--border2);
    border-radius: 10px;
    padding: 14px 18px;
    font-size: 13px;
    box-shadow: 0 8px 30px rgba(0,0,0,0.3);
    z-index: 9999;
    display: flex; align-items: center; gap: 10px;
    animation: toastIn 0.3s ease;
    max-width: 340px;
  }
  @keyframes toastIn { from { opacity:0; transform:translateY(10px); } to { opacity:1; transform:none; } }
  .toast.success { border-left: 3px solid var(--green); }
  .toast.error { border-left: 3px solid var(--red); }

  ::-webkit-scrollbar { width: 6px; }
  ::-webkit-scrollbar-track { background: transparent; }
  ::-webkit-scrollbar-thumb { background: var(--border2); border-radius: 3px; }
</style>
</head>
<body>
<div class="layout">
  <!-- Sidebar -->
  <nav class="sidebar">
    <div class="logo">
      <div class="logo-icon">⚡</div>
      <div class="logo-name">CronPanel</div>
      <div class="logo-sub">定时任务管理器</div>
    </div>
    <div class="nav">
      <div class="nav-label">功能</div>
      <div class="nav-item active" onclick="showPage('dashboard')">
        <span class="nav-icon">⊞</span> 仪表盘
      </div>
      <div class="nav-item" onclick="openAddModal()">
        <span class="nav-icon">＋</span> 添加任务
      </div>
    </div>
    <div class="sidebar-footer">
      <span class="status-dot"></span>
      服务运行中
    </div>
  </nav>

  <!-- Main -->
  <main class="main">
    <!-- Header -->
    <div class="page-header">
      <div>
        <div class="page-title">仪表盘</div>
        <div class="page-sub">管理您的 Linux Crontab 定时任务</div>
      </div>
      <div style="display:flex;gap:10px;align-items:center;">
        <button class="btn btn-ghost" onclick="loadJobs()">↺ 刷新</button>
        <button class="btn btn-primary" onclick="openAddModal()">＋ 添加任务</button>
      </div>
    </div>

    <!-- Stats -->
    <div class="stats-row">
      <div class="stat-card total">
        <div class="stat-label">全部任务</div>
        <div class="stat-value" id="stat-total">—</div>
        <div class="stat-sub">已配置的定时任务</div>
      </div>
      <div class="stat-card active">
        <div class="stat-label">运行中</div>
        <div class="stat-value" id="stat-active">—</div>
        <div class="stat-sub">当前启用的任务</div>
      </div>
      <div class="stat-card disabled">
        <div class="stat-label">已停用</div>
        <div class="stat-value" id="stat-disabled">—</div>
        <div class="stat-sub">已禁用的任务</div>
      </div>
    </div>

    <!-- Job List -->
    <div class="section-header">
      <div class="section-title">任务列表</div>
      <div id="last-refresh" style="font-size:11px;color:var(--text3)">正在加载...</div>
    </div>
    <div id="job-list" class="job-list">
      <div class="loading"><div class="spinner"></div>加载中...</div>
    </div>
  </main>
</div>

<!-- Add Job Modal -->
<div class="modal-overlay" id="add-modal">
  <div class="modal">
    <div class="modal-header">
      <div class="modal-title">添加定时任务</div>
      <button class="modal-close" onclick="closeModal()">✕</button>
    </div>
    <div class="modal-body">
      <!-- Mode tabs -->
      <div class="mode-tabs">
        <button class="mode-tab active" onclick="setMode('daily')">每天</button>
        <button class="mode-tab" onclick="setMode('interval')">每N天</button>
        <button class="mode-tab" onclick="setMode('weekly')">每周</button>
        <button class="mode-tab" onclick="setMode('monthly')">每月</button>
        <button class="mode-tab" onclick="setMode('custom')">自定义</button>
      </div>

      <!-- Time (shared) -->
      <div id="time-section">
        <div class="time-row">
          <div class="form-group">
            <label class="form-label">小时 (0-23)</label>
            <input type="number" class="form-input" id="f-hour" min="0" max="23" value="0" onchange="updatePreview()">
          </div>
          <div class="form-group">
            <label class="form-label">分钟 (0-59)</label>
            <input type="number" class="form-input" id="f-minute" min="0" max="59" value="0" onchange="updatePreview()">
          </div>
        </div>
      </div>

      <!-- Daily mode -->
      <div class="mode-section active" id="mode-daily">
        <div style="font-size:12px;color:var(--text2);margin-bottom:16px;padding:10px 14px;background:var(--surface2);border-radius:8px;">
          每天在指定时间运行
        </div>
      </div>

      <!-- Interval mode -->
      <div class="mode-section" id="mode-interval">
        <div class="form-group">
          <label class="form-label">每隔几天运行一次</label>
          <input type="number" class="form-input" id="f-days" min="1" max="30" value="1" placeholder="例如: 3" onchange="updatePreview()">
        </div>
      </div>

      <!-- Weekly mode -->
      <div class="mode-section" id="mode-weekly">
        <div class="form-group">
          <label class="form-label">星期几</label>
          <select class="form-select" id="f-weekday" onchange="updatePreview()">
            <option value="0">周日 (Sunday)</option>
            <option value="1">周一 (Monday)</option>
            <option value="2">周二 (Tuesday)</option>
            <option value="3">周三 (Wednesday)</option>
            <option value="4">周四 (Thursday)</option>
            <option value="5">周五 (Friday)</option>
            <option value="6">周六 (Saturday)</option>
          </select>
        </div>
      </div>

      <!-- Monthly mode -->
      <div class="mode-section" id="mode-monthly">
        <div class="time-row">
          <div class="form-group">
            <label class="form-label">每月第几天 (1-31)</label>
            <input type="number" class="form-input" id="f-monthday" min="1" max="31" value="1" onchange="updatePreview()">
          </div>
          <div class="form-group">
            <label class="form-label">月份 (* = 每月)</label>
            <select class="form-select" id="f-month" onchange="updatePreview()">
              <option value="*">每月</option>
              <option value="1">1月</option>
              <option value="2">2月</option>
              <option value="3">3月</option>
              <option value="4">4月</option>
              <option value="5">5月</option>
              <option value="6">6月</option>
              <option value="7">7月</option>
              <option value="8">8月</option>
              <option value="9">9月</option>
              <option value="10">10月</option>
              <option value="11">11月</option>
              <option value="12">12月</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Custom mode -->
      <div class="mode-section" id="mode-custom">
        <div class="form-group">
          <label class="form-label">Cron 表达式 (分 时 日 月 周)</label>
          <input type="text" class="form-input" id="f-custom" placeholder="0 2 * * 1" style="font-family:'DM Mono',monospace" onchange="updatePreview()" oninput="updatePreview()">
        </div>
      </div>

      <!-- Preview -->
      <div class="cron-preview">
        <div class="cron-preview-label">Cron 表达式预览</div>
        <span id="cron-preview-val">0 0 * * *</span>
      </div>

      <!-- Command section -->
      <div class="form-group">
        <label class="form-label">命令类型</label>
        <select class="form-select" id="f-cmd-type" onchange="updateCmdType()">
          <option value="cmd">直接命令</option>
          <option value="script-path">Shell 脚本路径</option>
          <option value="script-content">编写 Shell 脚本</option>
        </select>
      </div>

      <div id="cmd-section-cmd">
        <div class="form-group">
          <label class="form-label">执行命令</label>
          <input type="text" class="form-input" id="f-command" placeholder="例如: /usr/bin/python3 /home/user/script.py">
        </div>
      </div>
      <div id="cmd-section-script-path" style="display:none">
        <div class="form-group">
          <label class="form-label">Shell 脚本路径</label>
          <input type="text" class="form-input" id="f-script-path" placeholder="/home/user/backup.sh">
        </div>
      </div>
      <div id="cmd-section-script-content" style="display:none">
        <div class="form-group">
          <label class="form-label">Shell 脚本内容</label>
          <textarea class="form-textarea" id="f-script-content" placeholder="#!/bin/bash
echo 'Hello World'
date >> /tmp/cron.log"></textarea>
          <div style="font-size:11px;color:var(--text3);margin-top:6px;">脚本将保存到服务器并自动执行</div>
        </div>
      </div>

      <div class="form-group">
        <label class="form-label">备注说明（可选）</label>
        <input type="text" class="form-input" id="f-comment" placeholder="任务描述...">
      </div>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" onclick="closeModal()">取消</button>
      <button class="btn btn-primary" onclick="submitJob()">✓ 添加任务</button>
    </div>
  </div>
</div>

<script>
let currentMode = 'daily';
let jobs = [];

async function loadJobs() {
  document.getElementById('last-refresh').textContent = '正在刷新...';
  try {
    const res = await fetch('/api/jobs');
    const data = await res.json();
    if (data.success) {
      jobs = data.data || [];
      renderJobs(jobs);
      updateStats(jobs);
      document.getElementById('last-refresh').textContent = '更新于 ' + new Date().toLocaleTimeString('zh-CN');
    }
  } catch(e) {
    showToast('加载失败: ' + e.message, 'error');
  }
}

function renderJobs(jobs) {
  const el = document.getElementById('job-list');
  if (!jobs || jobs.length === 0) {
    el.innerHTML = '<div class="empty-state"><div class="empty-icon">⏰</div><div class="empty-text">暂无定时任务<br><span style="font-size:12px">点击右上角「添加任务」开始</span></div></div>';
    return;
  }
  el.innerHTML = jobs.map(job => {
    const isOn = job.enabled;
    const label = job.comment || job.command.substring(0, 40) + (job.command.length > 40 ? '...' : '');
    return '<div class="job-card ' + (isOn ? '' : 'disabled') + '">' +
      '<div class="job-status ' + (isOn ? 'on' : 'off') + '"></div>' +
      '<div class="job-info">' +
        '<div class="job-comment">' + escHtml(label) + '</div>' +
        '<div class="job-command">' + escHtml(job.command) + '</div>' +
      '</div>' +
      '<span class="job-schedule">' + escHtml(job.schedule) + '</span>' +
      '<div class="job-actions">' +
        '<button class="btn btn-ghost btn-sm" onclick="toggleJob(\'' + job.id + '\')" title="' + (isOn ? '停用' : '启用') + '">' + (isOn ? '⏸' : '▶') + '</button>' +
        '<button class="btn btn-danger btn-sm" onclick="deleteJob(\'' + job.id + '\')" title="删除">✕</button>' +
      '</div>' +
    '</div>';
  }).join('');
}

function updateStats(jobs) {
  const total = jobs.length;
  const active = jobs.filter(j => j.enabled).length;
  document.getElementById('stat-total').textContent = total;
  document.getElementById('stat-active').textContent = active;
  document.getElementById('stat-disabled').textContent = total - active;
}

function escHtml(s) {
  return s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
}

async function deleteJob(id) {
  if (!confirm('确认删除该任务？')) return;
  try {
    const res = await fetch('/api/jobs/delete', { method: 'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify({id}) });
    const data = await res.json();
    if (data.success) { showToast('任务已删除', 'success'); loadJobs(); }
    else showToast(data.message, 'error');
  } catch(e) { showToast(e.message, 'error'); }
}

async function toggleJob(id) {
  try {
    const res = await fetch('/api/jobs/toggle', { method: 'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify({id}) });
    const data = await res.json();
    if (data.success) { loadJobs(); }
    else showToast(data.message, 'error');
  } catch(e) { showToast(e.message, 'error'); }
}

function setMode(mode) {
  currentMode = mode;
  document.querySelectorAll('.mode-tab').forEach(t => t.classList.remove('active'));
  event.target.classList.add('active');
  document.querySelectorAll('.mode-section').forEach(s => s.classList.remove('active'));
  document.getElementById('mode-' + mode).classList.add('active');
  const showTime = mode !== 'custom';
  document.getElementById('time-section').style.display = showTime ? '' : 'none';
  updatePreview();
}

function getCronExpr() {
  const h = document.getElementById('f-hour').value || '0';
  const m = document.getElementById('f-minute').value || '0';
  if (currentMode === 'custom') return document.getElementById('f-custom').value || '0 0 * * *';
  if (currentMode === 'daily') return m + ' ' + h + ' * * *';
  if (currentMode === 'interval') {
    const d = document.getElementById('f-days').value || '1';
    return m + ' ' + h + ' */' + d + ' * *';
  }
  if (currentMode === 'weekly') {
    const wd = document.getElementById('f-weekday').value;
    return m + ' ' + h + ' * * ' + wd;
  }
  if (currentMode === 'monthly') {
    const md = document.getElementById('f-monthday').value || '1';
    const mo = document.getElementById('f-month').value || '*';
    return m + ' ' + h + ' ' + md + ' ' + mo + ' *';
  }
  return m + ' ' + h + ' * * *';
}

function updatePreview() {
  document.getElementById('cron-preview-val').textContent = getCronExpr();
}

function updateCmdType() {
  const t = document.getElementById('f-cmd-type').value;
  document.getElementById('cmd-section-cmd').style.display = t === 'cmd' ? '' : 'none';
  document.getElementById('cmd-section-script-path').style.display = t === 'script-path' ? '' : 'none';
  document.getElementById('cmd-section-script-content').style.display = t === 'script-content' ? '' : 'none';
}

async function submitJob() {
  const cmdType = document.getElementById('f-cmd-type').value;
  const body = {
    mode: currentMode,
    days: document.getElementById('f-days').value,
    weekday: document.getElementById('f-weekday').value,
    monthDay: document.getElementById('f-monthday').value,
    month: document.getElementById('f-month').value,
    hour: document.getElementById('f-hour').value,
    minute: document.getElementById('f-minute').value,
    comment: document.getElementById('f-comment').value,
    customCron: document.getElementById('f-custom').value,
    command: cmdType === 'cmd' ? document.getElementById('f-command').value : '',
    scriptPath: cmdType === 'script-path' ? document.getElementById('f-script-path').value : '',
    scriptContent: cmdType === 'script-content' ? document.getElementById('f-script-content').value : '',
  };
  try {
    const res = await fetch('/api/jobs/add', { method: 'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify(body) });
    const data = await res.json();
    if (data.success) {
      showToast('任务添加成功！', 'success');
      closeModal();
      loadJobs();
    } else {
      showToast(data.message || '添加失败', 'error');
    }
  } catch(e) { showToast(e.message, 'error'); }
}

function openAddModal() {
  document.getElementById('add-modal').classList.add('open');
  updatePreview();
}
function closeModal() {
  document.getElementById('add-modal').classList.remove('open');
}
document.getElementById('add-modal').addEventListener('click', function(e) {
  if (e.target === this) closeModal();
});

let toastTimer;
function showToast(msg, type) {
  const existing = document.querySelector('.toast');
  if (existing) existing.remove();
  clearTimeout(toastTimer);
  const div = document.createElement('div');
  div.className = 'toast ' + (type || '');
  div.innerHTML = (type === 'success' ? '✓ ' : '✕ ') + msg;
  document.body.appendChild(div);
  toastTimer = setTimeout(() => div.remove(), 3000);
}

function showPage(p) {}

loadJobs();
setInterval(loadJobs, 30000);
</script>
</body>
</html>`
