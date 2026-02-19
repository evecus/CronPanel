package main

var indexHTML = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>CronPanel — 定时任务管理</title>
<link rel="icon" type="image/svg+xml" href="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32'%3E%3Crect width='32' height='32' rx='8' fill='%234f46e5'/%3E%3Ccircle cx='16' cy='16' r='9' fill='none' stroke='white' stroke-width='2'/%3E%3Cline x1='16' y1='16' x2='16' y2='10' stroke='white' stroke-width='2.2' stroke-linecap='round'/%3E%3Cline x1='16' y1='16' x2='20' y2='18' stroke='%23a5b4fc' stroke-width='2.2' stroke-linecap='round'/%3E%3Ccircle cx='16' cy='16' r='2' fill='white'/%3E%3C/svg%3E">
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;500&family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap" rel="stylesheet">
<style>
  :root {
    --bg: #f1f4f9;
    --sidebar-bg: #1e1b4b;
    --sidebar-border: rgba(255,255,255,0.08);
    --surface: #ffffff;
    --surface2: #f8fafc;
    --border: #e2e8f0;
    --border2: #cbd5e1;
    --accent: #4f46e5;
    --accent2: #7c3aed;
    --accent-light: #eef2ff;
    --accent-text: #4338ca;
    --green: #059669;
    --green-bg: #ecfdf5;
    --red: #dc2626;
    --red-bg: #fef2f2;
    --text: #0f172a;
    --text2: #475569;
    --text3: #94a3b8;
    --sidebar-text: rgba(255,255,255,0.65);
    --shadow: 0 1px 3px rgba(0,0,0,0.06), 0 1px 2px rgba(0,0,0,0.04);
    --shadow-md: 0 4px 16px rgba(0,0,0,0.08);
    --shadow-lg: 0 20px 60px rgba(0,0,0,0.12), 0 4px 12px rgba(0,0,0,0.08);
  }
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body { background: var(--bg); color: var(--text); font-family: 'Plus Jakarta Sans', sans-serif; font-size: 14px; min-height: 100vh; -webkit-font-smoothing: antialiased; }
  .layout { display: flex; min-height: 100vh; }

  .sidebar { width: 240px; background: var(--sidebar-bg); display: flex; flex-direction: column; position: fixed; top: 0; left: 0; bottom: 0; z-index: 100; background-image: radial-gradient(ellipse at top left, rgba(99,102,241,0.3) 0%, transparent 60%), radial-gradient(ellipse at bottom, rgba(124,58,237,0.2) 0%, transparent 60%); }
  .logo { padding: 26px 22px 22px; border-bottom: 1px solid var(--sidebar-border); display: flex; align-items: center; gap: 12px; }
  .logo-icon { width: 36px; height: 36px; flex-shrink: 0; }
  .logo-icon svg { width: 100%; height: 100%; }
  .logo-name { font-weight: 800; font-size: 15px; color: white; letter-spacing: 0.2px; line-height: 1.2; }
  .logo-sub { font-size: 11px; color: var(--sidebar-text); margin-top: 1px; }
  .nav { padding: 18px 12px; flex: 1; }
  .nav-label { font-size: 10px; font-weight: 700; text-transform: uppercase; letter-spacing: 1.5px; color: rgba(255,255,255,0.3); padding: 0 10px; margin-bottom: 7px; }
  .nav-item { display: flex; align-items: center; gap: 10px; padding: 9px 12px; border-radius: 10px; cursor: pointer; font-size: 13.5px; font-weight: 500; color: var(--sidebar-text); transition: all 0.15s; margin-bottom: 2px; }
  .nav-item:hover { background: rgba(255,255,255,0.08); color: white; }
  .nav-item.active { background: rgba(255,255,255,0.14); color: white; }
  .nav-icon { width: 26px; height: 26px; border-radius: 7px; display: flex; align-items: center; justify-content: center; font-size: 13px; flex-shrink: 0; }
  .nav-item.active .nav-icon { background: rgba(255,255,255,0.18); }
  .sidebar-footer { padding: 14px 22px; border-top: 1px solid var(--sidebar-border); }
  .status-badge { display: inline-flex; align-items: center; gap: 6px; font-size: 11.5px; color: var(--sidebar-text); }
  .status-dot { width: 6px; height: 6px; background: #34d399; border-radius: 50%; box-shadow: 0 0 6px #34d399; animation: pulse 2.5s ease-in-out infinite; }
  @keyframes pulse { 0%,100%{opacity:1;transform:scale(1)} 50%{opacity:0.7;transform:scale(0.9)} }

  .main { margin-left: 240px; flex: 1; padding: 36px 40px; max-width: 1080px; }
  .page-header { display: flex; align-items: flex-start; justify-content: space-between; margin-bottom: 28px; gap: 16px; }
  .page-title { font-size: 26px; font-weight: 800; color: var(--text); letter-spacing: -0.5px; line-height: 1.2; }
  .page-sub { font-size: 13px; color: var(--text3); margin-top: 4px; }
  .btn { display: inline-flex; align-items: center; gap: 7px; padding: 9px 18px; border-radius: 10px; border: none; cursor: pointer; font-size: 13.5px; font-weight: 600; font-family: 'Plus Jakarta Sans', sans-serif; transition: all 0.15s; white-space: nowrap; }
  .btn-primary { background: var(--accent); color: white; box-shadow: 0 2px 8px rgba(79,70,229,0.3); }
  .btn-primary:hover { background: #4338ca; transform: translateY(-1px); box-shadow: 0 4px 16px rgba(79,70,229,0.35); }
  .btn-primary:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }
  .btn-ghost { background: white; color: var(--text2); border: 1.5px solid var(--border); box-shadow: var(--shadow); }
  .btn-ghost:hover { background: var(--surface2); color: var(--text); border-color: var(--border2); }
  .btn-danger { background: var(--red-bg); color: var(--red); border: 1.5px solid #fecaca; }
  .btn-danger:hover { background: #fee2e2; }
  .btn-sm { padding: 5px 11px; font-size: 12px; border-radius: 7px; }

  .stats-row { display: grid; grid-template-columns: repeat(3, 1fr); gap: 14px; margin-bottom: 26px; }
  .stat-card { background: white; border: 1.5px solid var(--border); border-radius: 14px; padding: 18px 20px; position: relative; overflow: hidden; box-shadow: var(--shadow); }
  .stat-card-accent { position: absolute; top: 0; left: 0; width: 4px; height: 100%; }
  .stat-card.total .stat-card-accent { background: linear-gradient(180deg, #4f46e5, #7c3aed); }
  .stat-card.active .stat-card-accent { background: linear-gradient(180deg, #10b981, #059669); }
  .stat-card.disabled .stat-card-accent { background: linear-gradient(180deg, #94a3b8, #64748b); }
  .stat-label { font-size: 11.5px; font-weight: 600; color: var(--text3); text-transform: uppercase; letter-spacing: 0.8px; }
  .stat-value { font-size: 34px; font-weight: 800; margin: 6px 0 2px; color: var(--text); letter-spacing: -1px; line-height: 1; }
  .stat-sub { font-size: 12px; color: var(--text3); }

  .section-header { display: flex; align-items: center; justify-content: space-between; margin-bottom: 12px; }
  .section-title { font-size: 15px; font-weight: 700; color: var(--text); }
  .job-list { display: flex; flex-direction: column; gap: 8px; }
  .job-card { background: white; border: 1.5px solid var(--border); border-radius: 12px; padding: 14px 18px; display: flex; align-items: center; gap: 14px; transition: all 0.15s; box-shadow: var(--shadow); animation: slideIn 0.22s ease; }
  @keyframes slideIn { from { opacity:0; transform:translateY(5px); } to { opacity:1; transform:none; } }
  .job-card:hover { border-color: var(--border2); box-shadow: var(--shadow-md); transform: translateY(-1px); }
  .job-card.disabled { opacity: 0.6; }
  .job-status { width: 8px; height: 8px; border-radius: 50%; flex-shrink: 0; }
  .job-status.on { background: var(--green); box-shadow: 0 0 6px rgba(5,150,105,0.5); }
  .job-status.off { background: var(--text3); }
  .job-info { flex: 1; min-width: 0; }
  .job-comment { font-size: 13.5px; font-weight: 600; color: var(--text); margin-bottom: 5px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
  .job-command { font-family: 'IBM Plex Mono', monospace; font-size: 11px; color: #4f46e5; background: var(--accent-light); border: 1px solid #c7d2fe; border-radius: 5px; padding: 2px 8px; display: inline-block; max-width: 100%; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
  .job-schedule { font-family: 'IBM Plex Mono', monospace; font-size: 11.5px; color: #7c3aed; background: #f5f3ff; border: 1px solid #ddd6fe; padding: 4px 10px; border-radius: 20px; white-space: nowrap; flex-shrink: 0; font-weight: 500; }
  .job-actions { display: flex; gap: 6px; flex-shrink: 0; }

  .empty-state { text-align: center; padding: 60px 20px; background: white; border: 1.5px dashed var(--border2); border-radius: 16px; }
  .empty-icon { width: 60px; height: 60px; margin: 0 auto 14px; background: var(--accent-light); border-radius: 14px; display: flex; align-items: center; justify-content: center; font-size: 26px; }
  .empty-title { font-size: 15px; font-weight: 600; color: var(--text2); margin-bottom: 5px; }
  .empty-sub { font-size: 13px; color: var(--text3); }

  .modal-overlay { position: fixed; inset: 0; background: rgba(15,23,42,0.55); backdrop-filter: blur(6px); z-index: 1000; display: none; align-items: center; justify-content: center; padding: 20px; }
  .modal-overlay.open { display: flex; }
  .modal { background: white; border: 1.5px solid var(--border); border-radius: 20px; width: 100%; max-width: 575px; max-height: 90vh; overflow-y: auto; box-shadow: var(--shadow-lg); animation: modalIn 0.22s cubic-bezier(0.34,1.56,0.64,1); }
  @keyframes modalIn { from { opacity:0; transform:scale(0.94) translateY(10px); } to { opacity:1; transform:none; } }
  .modal-header { padding: 20px 22px 16px; border-bottom: 1.5px solid var(--border); display: flex; align-items: center; justify-content: space-between; position: sticky; top: 0; background: white; border-radius: 20px 20px 0 0; z-index: 1; }
  .modal-title-wrap { display: flex; align-items: center; gap: 10px; }
  .modal-title-icon { width: 30px; height: 30px; background: var(--accent-light); border-radius: 8px; display: flex; align-items: center; justify-content: center; font-size: 15px; }
  .modal-title { font-size: 15.5px; font-weight: 700; color: var(--text); }
  .modal-close { background: none; border: none; cursor: pointer; color: var(--text3); font-size: 17px; width: 30px; height: 30px; border-radius: 7px; display: flex; align-items: center; justify-content: center; transition: all 0.12s; }
  .modal-close:hover { background: var(--surface2); color: var(--text); }
  .modal-body { padding: 20px 22px; }
  .modal-footer { padding: 14px 22px; border-top: 1.5px solid var(--border); display: flex; gap: 10px; justify-content: flex-end; background: var(--surface2); border-radius: 0 0 20px 20px; }

  .form-group { margin-bottom: 16px; }
  .form-label { display: flex; align-items: center; gap: 5px; font-size: 11.5px; font-weight: 700; color: var(--text2); margin-bottom: 6px; text-transform: uppercase; letter-spacing: 0.7px; }
  .form-input, .form-select, .form-textarea { width: 100%; background: white; border: 1.5px solid var(--border); border-radius: 9px; padding: 9px 13px; color: var(--text); font-size: 13.5px; font-family: 'Plus Jakarta Sans', sans-serif; outline: none; transition: border-color 0.15s, box-shadow 0.15s; }
  .form-input:focus, .form-select:focus, .form-textarea:focus { border-color: var(--accent); box-shadow: 0 0 0 3px rgba(79,70,229,0.1); }
  .form-input::placeholder { color: var(--text3); }
  .form-select { cursor: pointer; appearance: none; background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='8' fill='none'%3E%3Cpath d='M1 1l5 5 5-5' stroke='%2394a3b8' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E"); background-repeat: no-repeat; background-position: right 12px center; padding-right: 34px; }
  .form-textarea { min-height: 100px; resize: vertical; font-family: 'IBM Plex Mono', monospace; font-size: 12px; line-height: 1.6; }

  .mode-tabs { display: flex; background: var(--surface2); border: 1.5px solid var(--border); border-radius: 12px; padding: 4px; gap: 3px; margin-bottom: 18px; }
  .mode-tab { flex: 1; padding: 8px 4px; border-radius: 8px; border: none; background: none; color: var(--text3); font-size: 12px; font-weight: 600; cursor: pointer; transition: all 0.15s; font-family: 'Plus Jakarta Sans', sans-serif; text-align: center; white-space: nowrap; }
  .mode-tab:hover { color: var(--text2); }
  .mode-tab.active { background: white; color: var(--accent); box-shadow: var(--shadow); border: 1px solid var(--border); }
  .time-row { display: flex; gap: 12px; }
  .time-row .form-group { flex: 1; }
  .mode-section { display: none; }
  .mode-section.active { display: block; }
  .cron-preview { background: var(--accent-light); border: 1.5px solid #c7d2fe; border-radius: 10px; padding: 10px 14px; margin-bottom: 16px; display: flex; align-items: center; gap: 10px; }
  .cron-preview-label { font-size: 10px; font-weight: 700; color: var(--accent-text); text-transform: uppercase; letter-spacing: 0.7px; white-space: nowrap; }
  .cron-preview-val { font-family: 'IBM Plex Mono', monospace; font-size: 13px; color: var(--accent-text); font-weight: 500; }
  .info-box { font-size: 12.5px; color: var(--text2); padding: 10px 13px; background: var(--surface2); border: 1px solid var(--border); border-radius: 8px; margin-bottom: 14px; }
  .divider { height: 1px; background: var(--border); margin: 16px 0; }
  .section-sub-label { font-size: 11.5px; font-weight: 700; text-transform: uppercase; letter-spacing: 0.7px; color: var(--text3); margin-bottom: 12px; }

  .loading { text-align: center; padding: 48px; color: var(--text3); }
  .spinner { width: 26px; height: 26px; border: 2.5px solid var(--border); border-top-color: var(--accent); border-radius: 50%; animation: spin 0.7s linear infinite; margin: 0 auto 12px; }
  @keyframes spin { to { transform: rotate(360deg); } }

  .toast { position: fixed; bottom: 22px; right: 22px; background: white; border: 1.5px solid var(--border); border-radius: 12px; padding: 12px 18px; font-size: 13.5px; font-weight: 500; box-shadow: var(--shadow-lg); z-index: 9999; display: flex; align-items: center; gap: 9px; animation: toastIn 0.25s cubic-bezier(0.34,1.56,0.64,1); max-width: 320px; color: var(--text); }
  @keyframes toastIn { from { opacity:0; transform:translateY(6px) scale(0.95); } to { opacity:1; transform:none; } }
  .toast.success { border-left: 3.5px solid var(--green); }
  .toast.error { border-left: 3.5px solid var(--red); }
  ::-webkit-scrollbar { width: 5px; }
  ::-webkit-scrollbar-track { background: transparent; }
  ::-webkit-scrollbar-thumb { background: var(--border2); border-radius: 3px; }
</style>
</head>
<body>
<div class="layout">
  <nav class="sidebar">
    <div class="logo">
      <div class="logo-icon">
        <svg viewBox="0 0 36 36" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect width="36" height="36" rx="9" fill="white" fill-opacity="0.15"/>
          <circle cx="18" cy="18" r="10" stroke="white" stroke-width="2"/>
          <line x1="18" y1="18" x2="18" y2="11" stroke="white" stroke-width="2.2" stroke-linecap="round"/>
          <line x1="18" y1="18" x2="23" y2="21" stroke="#a5b4fc" stroke-width="2.2" stroke-linecap="round"/>
          <circle cx="18" cy="18" r="2" fill="white"/>
          <circle cx="18" cy="8.5" r="1.2" fill="white" opacity="0.5"/>
          <circle cx="27.5" cy="18" r="1.2" fill="white" opacity="0.5"/>
          <circle cx="18" cy="27.5" r="1.2" fill="white" opacity="0.5"/>
          <circle cx="8.5" cy="18" r="1.2" fill="white" opacity="0.5"/>
        </svg>
      </div>
      <div>
        <div class="logo-name">CronPanel</div>
        <div class="logo-sub">定时任务管理器</div>
      </div>
    </div>
    <div class="nav">
      <div class="nav-label">功能</div>
      <div class="nav-item active" onclick="showPage('dashboard')">
        <div class="nav-icon">⊞</div> 仪表盘
      </div>
      <div class="nav-item" onclick="openAddModal()">
        <div class="nav-icon">＋</div> 添加任务
      </div>
    </div>
    <div class="sidebar-footer">
      <div class="status-badge"><div class="status-dot"></div>服务运行中</div>
    </div>
  </nav>

  <main class="main">
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

    <div class="stats-row">
      <div class="stat-card total">
        <div class="stat-card-accent"></div>
        <div class="stat-label">全部任务</div>
        <div class="stat-value" id="stat-total">—</div>
        <div class="stat-sub">已配置的定时任务</div>
      </div>
      <div class="stat-card active">
        <div class="stat-card-accent"></div>
        <div class="stat-label">运行中</div>
        <div class="stat-value" id="stat-active">—</div>
        <div class="stat-sub">当前启用的任务</div>
      </div>
      <div class="stat-card disabled">
        <div class="stat-card-accent"></div>
        <div class="stat-label">已停用</div>
        <div class="stat-value" id="stat-disabled">—</div>
        <div class="stat-sub">已禁用的任务</div>
      </div>
    </div>

    <div class="section-header">
      <div class="section-title">任务列表</div>
      <div id="last-refresh" style="font-size:11.5px;color:var(--text3)">正在加载...</div>
    </div>
    <div id="job-list" class="job-list">
      <div class="loading"><div class="spinner"></div>加载中...</div>
    </div>
  </main>
</div>

<div class="modal-overlay" id="add-modal">
  <div class="modal">
    <div class="modal-header">
      <div class="modal-title-wrap">
        <div class="modal-title-icon">⏰</div>
        <div class="modal-title">添加定时任务</div>
      </div>
      <button class="modal-close" onclick="closeModal()">✕</button>
    </div>
    <div class="modal-body">
      <div class="section-sub-label">执行计划</div>
      <div class="mode-tabs">
        <button class="mode-tab active" data-mode="daily" onclick="setMode('daily',this)">每天</button>
        <button class="mode-tab" data-mode="interval" onclick="setMode('interval',this)">每N天</button>
        <button class="mode-tab" data-mode="weekly" onclick="setMode('weekly',this)">每周</button>
        <button class="mode-tab" data-mode="monthly" onclick="setMode('monthly',this)">每月</button>
        <button class="mode-tab" data-mode="custom" onclick="setMode('custom',this)">自定义</button>
      </div>

      <div id="time-section">
        <div class="time-row">
          <div class="form-group">
            <label class="form-label">小时 (0-23)</label>
            <input type="number" class="form-input" id="f-hour" min="0" max="23" value="0" oninput="updatePreview()">
          </div>
          <div class="form-group">
            <label class="form-label">分钟 (0-59)</label>
            <input type="number" class="form-input" id="f-minute" min="0" max="59" value="0" oninput="updatePreview()">
          </div>
        </div>
      </div>

      <div class="mode-section active" id="mode-daily">
        <div class="info-box">每天在指定时间运行一次</div>
      </div>
      <div class="mode-section" id="mode-interval">
        <div class="form-group">
          <label class="form-label">每隔几天运行一次</label>
          <input type="number" class="form-input" id="f-days" min="1" max="30" value="2" placeholder="例如: 3" oninput="updatePreview()">
        </div>
      </div>
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
      <div class="mode-section" id="mode-monthly">
        <div class="time-row">
          <div class="form-group">
            <label class="form-label">每月第几天 (1-31)</label>
            <input type="number" class="form-input" id="f-monthday" min="1" max="31" value="1" oninput="updatePreview()">
          </div>
          <div class="form-group">
            <label class="form-label">月份</label>
            <select class="form-select" id="f-month" onchange="updatePreview()">
              <option value="*">每月</option>
              <option value="1">1月</option><option value="2">2月</option>
              <option value="3">3月</option><option value="4">4月</option>
              <option value="5">5月</option><option value="6">6月</option>
              <option value="7">7月</option><option value="8">8月</option>
              <option value="9">9月</option><option value="10">10月</option>
              <option value="11">11月</option><option value="12">12月</option>
            </select>
          </div>
        </div>
      </div>
      <div class="mode-section" id="mode-custom">
        <div class="form-group">
          <label class="form-label">Cron 表达式 <span style="font-size:10px;color:var(--text3);font-weight:400;text-transform:none;letter-spacing:0">(分 时 日 月 周)</span></label>
          <input type="text" class="form-input" id="f-custom" placeholder="例如: 0 2 * * 1" style="font-family:'IBM Plex Mono',monospace;font-size:13px" oninput="updatePreview()">
        </div>
      </div>

      <div class="cron-preview">
        <div class="cron-preview-label">CRON</div>
        <div class="cron-preview-val" id="cron-preview-val">0 0 * * *</div>
      </div>

      <div class="divider"></div>
      <div class="section-sub-label">执行命令</div>

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
          <input type="text" class="form-input" id="f-command" placeholder="例如: /usr/bin/python3 /home/user/script.py" style="font-family:'IBM Plex Mono',monospace;font-size:12.5px">
        </div>
      </div>
      <div id="cmd-section-script-path" style="display:none">
        <div class="form-group">
          <label class="form-label">Shell 脚本路径</label>
          <input type="text" class="form-input" id="f-script-path" placeholder="/home/user/backup.sh" style="font-family:'IBM Plex Mono',monospace;font-size:12.5px">
        </div>
      </div>
      <div id="cmd-section-script-content" style="display:none">
        <div class="form-group">
          <label class="form-label">Shell 脚本内容</label>
          <textarea class="form-textarea" id="f-script-content" placeholder="#!/bin/bash&#10;echo 'Hello World'&#10;date >> /tmp/cron.log"></textarea>
          <div style="font-size:11.5px;color:var(--text3);margin-top:5px;">脚本将保存到服务器并自动执行</div>
        </div>
      </div>
      <div class="form-group">
        <label class="form-label">备注说明 <span style="font-size:10px;font-weight:400;text-transform:none;letter-spacing:0;color:var(--text3)">(可选)</span></label>
        <input type="text" class="form-input" id="f-comment" placeholder="任务描述，便于识别...">
      </div>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" onclick="closeModal()">取消</button>
      <button class="btn btn-primary" id="submit-btn" onclick="submitJob()">✓ 添加任务</button>
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
    } else {
      showToast('加载失败: ' + (data.message || '未知错误'), 'error');
    }
  } catch(e) {
    showToast('连接失败: ' + e.message, 'error');
    document.getElementById('last-refresh').textContent = '加载失败';
  }
}

function renderJobs(jobs) {
  const el = document.getElementById('job-list');
  if (!jobs || jobs.length === 0) {
    el.innerHTML = '<div class="empty-state"><div class="empty-icon">⏰</div><div class="empty-title">暂无定时任务</div><div class="empty-sub">点击右上角「添加任务」开始</div></div>';
    return;
  }
  el.innerHTML = jobs.map(job => {
    const isOn = job.enabled;
    const label = job.comment || (job.command.length > 50 ? job.command.substring(0,50)+'...' : job.command);
    const id = escAttr(job.id);
    return '<div class="job-card ' + (isOn ? '' : 'disabled') + '">' +
      '<div class="job-status ' + (isOn ? 'on' : 'off') + '"></div>' +
      '<div class="job-info">' +
        '<div class="job-comment">' + escHtml(label) + '</div>' +
        '<span class="job-command">' + escHtml(job.command) + '</span>' +
      '</div>' +
      '<span class="job-schedule">' + escHtml(job.schedule) + '</span>' +
      '<div class="job-actions">' +
        '<button class="btn btn-ghost btn-sm" onclick="toggleJob(\'' + id + '\')">' + (isOn ? '⏸ 停用' : '▶ 启用') + '</button>' +
        '<button class="btn btn-danger btn-sm" onclick="deleteJob(\'' + id + '\')">✕</button>' +
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
  if (typeof s !== 'string') s = String(s||''
  );
  return s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
}
function escAttr(s) {
  if (typeof s !== 'string') s = String(s||''
  );
  return s.replace(/['"\\]/g, function(c){return '\\'+ c;});
}

async function deleteJob(id) {
  if (!confirm('确认删除该任务？此操作不可撤销。')) return;
  try {
    const res = await fetch('/api/jobs/delete', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id})});
    const data = await res.json();
    if (data.success) { showToast('任务已删除', 'success'); loadJobs(); }
    else showToast(data.message || '删除失败', 'error');
  } catch(e) { showToast(e.message, 'error'); }
}

async function toggleJob(id) {
  try {
    const res = await fetch('/api/jobs/toggle', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id})});
    const data = await res.json();
    if (data.success) { loadJobs(); }
    else showToast(data.message || '操作失败', 'error');
  } catch(e) { showToast(e.message, 'error'); }
}

function setMode(mode, btn) {
  currentMode = mode;
  document.querySelectorAll('.mode-tab').forEach(t => t.classList.remove('active'));
  if (btn) btn.classList.add('active');
  document.querySelectorAll('.mode-section').forEach(s => s.classList.remove('active'));
  const section = document.getElementById('mode-' + mode);
  if (section) section.classList.add('active');
  document.getElementById('time-section').style.display = (mode === 'custom') ? 'none' : '';
  updatePreview();
}

function getCronExpr() {
  const h = (document.getElementById('f-hour').value || '0').trim();
  const m = (document.getElementById('f-minute').value || '0').trim();
  if (currentMode === 'custom') return (document.getElementById('f-custom').value || '').trim() || '0 0 * * *';
  if (currentMode === 'daily') return m + ' ' + h + ' * * *';
  if (currentMode === 'interval') { const d = parseInt(document.getElementById('f-days').value)||1; return m + ' ' + h + ' */' + d + ' * *'; }
  if (currentMode === 'weekly') { const wd = document.getElementById('f-weekday').value; return m + ' ' + h + ' * * ' + wd; }
  if (currentMode === 'monthly') { const md = document.getElementById('f-monthday').value||'1'; const mo = document.getElementById('f-month').value||'*'; return m + ' ' + h + ' ' + md + ' ' + mo + ' *'; }
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
  let command = '', scriptPath = '', scriptContent = '';
  if (cmdType === 'cmd') {
    command = document.getElementById('f-command').value.trim();
    if (!command) { showToast('请输入执行命令', 'error'); document.getElementById('f-command').focus(); return; }
  } else if (cmdType === 'script-path') {
    scriptPath = document.getElementById('f-script-path').value.trim();
    if (!scriptPath) { showToast('请输入脚本路径', 'error'); document.getElementById('f-script-path').focus(); return; }
  } else if (cmdType === 'script-content') {
    scriptContent = document.getElementById('f-script-content').value.trim();
    if (!scriptContent) { showToast('请输入脚本内容', 'error'); document.getElementById('f-script-content').focus(); return; }
  }
  if (currentMode === 'custom') {
    const cron = document.getElementById('f-custom').value.trim();
    if (!cron) { showToast('请输入 Cron 表达式', 'error'); document.getElementById('f-custom').focus(); return; }
    if (cron.split(/\s+/).length !== 5) { showToast('Cron 格式错误，需要 5 个字段 (分 时 日 月 周)', 'error'); return; }
  }
  const body = {
    mode: currentMode,
    days: document.getElementById('f-days').value || '1',
    weekday: document.getElementById('f-weekday').value || '0',
    monthDay: document.getElementById('f-monthday').value || '1',
    month: document.getElementById('f-month').value || '*',
    hour: document.getElementById('f-hour').value || '0',
    minute: document.getElementById('f-minute').value || '0',
    comment: document.getElementById('f-comment').value.trim(),
    customCron: document.getElementById('f-custom').value.trim(),
    command, scriptPath, scriptContent,
  };
  const btn = document.getElementById('submit-btn');
  btn.disabled = true; btn.textContent = '提交中...';
  try {
    const res = await fetch('/api/jobs/add', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});
    const data = await res.json();
    if (data.success) { showToast('任务添加成功！', 'success'); closeModal(); resetForm(); loadJobs(); }
    else showToast(data.message || '添加失败，请检查参数', 'error');
  } catch(e) { showToast('网络错误: ' + e.message, 'error'); }
  finally { btn.disabled = false; btn.textContent = '✓ 添加任务'; }
}

function resetForm() {
  ['f-hour','f-minute','f-command','f-script-path','f-script-content','f-comment','f-custom'].forEach(id => {
    const el = document.getElementById(id);
    if (el) el.value = id === 'f-hour' || id === 'f-minute' ? '0' : (id === 'f-days' ? '2' : '');
  });
  document.getElementById('f-days').value = '2';
  document.getElementById('f-monthday').value = '1';
  document.getElementById('f-month').value = '*';
  document.getElementById('f-weekday').value = '0';
  document.getElementById('f-cmd-type').value = 'cmd';
  updateCmdType();
  const dailyBtn = document.querySelector('.mode-tab[data-mode="daily"]');
  setMode('daily', dailyBtn);
}

function openAddModal() { document.getElementById('add-modal').classList.add('open'); updatePreview(); }
function closeModal() { document.getElementById('add-modal').classList.remove('open'); }
document.getElementById('add-modal').addEventListener('click', function(e) { if (e.target === this) closeModal(); });

let toastTimer;
function showToast(msg, type) {
  const existing = document.querySelector('.toast');
  if (existing) existing.remove();
  clearTimeout(toastTimer);
  const div = document.createElement('div');
  div.className = 'toast ' + (type || '');
  div.innerHTML = '<span>' + (type === 'success' ? '✓' : '✕') + '</span>' + escHtml(msg);
  document.body.appendChild(div);
  toastTimer = setTimeout(() => { div.style.opacity = '0'; div.style.transition = 'opacity 0.3s'; setTimeout(() => div.remove(), 300); }, 3000);
}

function showPage(p) {}
loadJobs();
setInterval(loadJobs, 30000);
</script>
</body>
</html>`
