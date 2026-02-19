package main

var indexHTML = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>CronPanel</title>
<link rel="icon" type="image/svg+xml" href="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32'%3E%3Crect width='32' height='32' rx='8' fill='%234f46e5'/%3E%3Ccircle cx='16' cy='16' r='9' fill='none' stroke='white' stroke-width='2'/%3E%3Cline x1='16' y1='16' x2='16' y2='10' stroke='white' stroke-width='2.2' stroke-linecap='round'/%3E%3Cline x1='16' y1='16' x2='20' y2='18' stroke='%23a5b4fc' stroke-width='2.2' stroke-linecap='round'/%3E%3Ccircle cx='16' cy='16' r='2' fill='white'/%3E%3C/svg%3E">
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;500&family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap" rel="stylesheet">
<style>
  :root {
    --bg:#f1f4f9;--sidebar-bg:#1e1b4b;--sidebar-border:rgba(255,255,255,0.08);
    --surface:#fff;--surface2:#f8fafc;--border:#e2e8f0;--border2:#cbd5e1;
    --accent:#4f46e5;--accent2:#7c3aed;--accent-light:#eef2ff;--accent-text:#4338ca;
    --green:#059669;--red:#dc2626;--red-bg:#fef2f2;
    --text:#0f172a;--text2:#475569;--text3:#94a3b8;--sidebar-text:rgba(255,255,255,0.65);
    --shadow:0 1px 3px rgba(0,0,0,0.06),0 1px 2px rgba(0,0,0,0.04);
    --shadow-md:0 4px 16px rgba(0,0,0,0.08);
    --shadow-lg:0 20px 60px rgba(0,0,0,0.12),0 4px 12px rgba(0,0,0,0.08);
  }
  *{box-sizing:border-box;margin:0;padding:0}
  body{background:var(--bg);color:var(--text);font-family:'Plus Jakarta Sans',sans-serif;font-size:14px;min-height:100vh;-webkit-font-smoothing:antialiased}
  .layout{display:flex;min-height:100vh}

  /* SIDEBAR */
  .sidebar{width:240px;background:var(--sidebar-bg);display:flex;flex-direction:column;position:fixed;top:0;left:0;bottom:0;z-index:100;background-image:radial-gradient(ellipse at top left,rgba(99,102,241,0.3) 0%,transparent 60%),radial-gradient(ellipse at bottom,rgba(124,58,237,0.2) 0%,transparent 60%)}
  .logo{padding:26px 22px 22px;border-bottom:1px solid var(--sidebar-border);display:flex;align-items:center;gap:12px}
  .logo-icon{width:36px;height:36px;flex-shrink:0}
  .logo-icon svg{width:100%;height:100%}
  .logo-name{font-weight:800;font-size:15px;color:white;letter-spacing:0.2px;line-height:1.2}
  .logo-sub{font-size:11px;color:var(--sidebar-text);margin-top:1px}
  .nav{padding:18px 12px;flex:1}
  .nav-label{font-size:10px;font-weight:700;text-transform:uppercase;letter-spacing:1.5px;color:rgba(255,255,255,0.3);padding:0 10px;margin-bottom:7px}
  .nav-item{display:flex;align-items:center;gap:10px;padding:9px 12px;border-radius:10px;cursor:pointer;font-size:13.5px;font-weight:500;color:var(--sidebar-text);transition:all 0.15s;margin-bottom:2px}
  .nav-item:hover{background:rgba(255,255,255,0.08);color:white}
  .nav-item.active{background:rgba(255,255,255,0.14);color:white}
  .nav-icon{width:26px;height:26px;border-radius:7px;display:flex;align-items:center;justify-content:center;font-size:13px;flex-shrink:0}
  .nav-item.active .nav-icon{background:rgba(255,255,255,0.18)}
  .sidebar-footer{padding:14px 22px;border-top:1px solid var(--sidebar-border);display:flex;align-items:center;justify-content:space-between}
  .status-badge{display:inline-flex;align-items:center;gap:6px;font-size:11.5px;color:var(--sidebar-text)}
  .status-dot{width:6px;height:6px;background:#34d399;border-radius:50%;box-shadow:0 0 6px #34d399;animation:pulse 2.5s ease-in-out infinite}
  @keyframes pulse{0%,100%{opacity:1;transform:scale(1)}50%{opacity:0.7;transform:scale(0.9)}}
  .lang-btn{background:rgba(255,255,255,0.1);border:none;color:rgba(255,255,255,0.7);font-size:11px;font-weight:700;padding:4px 8px;border-radius:6px;cursor:pointer;font-family:'Plus Jakarta Sans',sans-serif;letter-spacing:0.5px;transition:all 0.15s}
  .lang-btn:hover{background:rgba(255,255,255,0.18);color:white}

  /* MAIN */
  .main{margin-left:240px;flex:1;padding:36px 40px;max-width:1080px}
  .page-header{display:flex;align-items:flex-start;justify-content:space-between;margin-bottom:28px;gap:16px}
  .page-title{font-size:26px;font-weight:800;color:var(--text);letter-spacing:-0.5px;line-height:1.2}
  .page-sub{font-size:13px;color:var(--text3);margin-top:4px}
  .btn{display:inline-flex;align-items:center;gap:7px;padding:9px 18px;border-radius:10px;border:none;cursor:pointer;font-size:13.5px;font-weight:600;font-family:'Plus Jakarta Sans',sans-serif;transition:all 0.15s;white-space:nowrap}
  .btn-primary{background:var(--accent);color:white;box-shadow:0 2px 8px rgba(79,70,229,0.3)}
  .btn-primary:hover{background:#4338ca;transform:translateY(-1px);box-shadow:0 4px 16px rgba(79,70,229,0.35)}
  .btn-primary:disabled{opacity:0.6;cursor:not-allowed;transform:none}
  .btn-ghost{background:white;color:var(--text2);border:1.5px solid var(--border);box-shadow:var(--shadow)}
  .btn-ghost:hover{background:var(--surface2);color:var(--text);border-color:var(--border2)}
  .btn-danger{background:var(--red-bg);color:var(--red);border:1.5px solid #fecaca}
  .btn-danger:hover{background:#fee2e2}
  .btn-warning{background:#fffbeb;color:#d97706;border:1.5px solid #fde68a}
  .btn-warning:hover{background:#fef3c7}
  .btn-sm{padding:5px 11px;font-size:12px;border-radius:7px}

  /* STATS */
  .stats-row{display:grid;grid-template-columns:repeat(3,1fr);gap:14px;margin-bottom:26px}
  .stat-card{background:white;border:1.5px solid var(--border);border-radius:14px;padding:18px 20px;position:relative;overflow:hidden;box-shadow:var(--shadow)}
  .stat-card-accent{position:absolute;top:0;left:0;width:4px;height:100%}
  .stat-card.total .stat-card-accent{background:linear-gradient(180deg,#4f46e5,#7c3aed)}
  .stat-card.active .stat-card-accent{background:linear-gradient(180deg,#10b981,#059669)}
  .stat-card.disabled .stat-card-accent{background:linear-gradient(180deg,#94a3b8,#64748b)}
  .stat-label{font-size:11.5px;font-weight:600;color:var(--text3);text-transform:uppercase;letter-spacing:0.8px}
  .stat-value{font-size:34px;font-weight:800;margin:6px 0 2px;color:var(--text);letter-spacing:-1px;line-height:1}
  .stat-sub{font-size:12px;color:var(--text3)}

  /* JOB LIST */
  .section-header{display:flex;align-items:center;justify-content:space-between;margin-bottom:12px}
  .section-title{font-size:15px;font-weight:700;color:var(--text)}
  .job-list{display:flex;flex-direction:column;gap:8px}
  .job-card{background:white;border:1.5px solid var(--border);border-radius:12px;padding:14px 18px;display:flex;align-items:center;gap:14px;transition:all 0.15s;box-shadow:var(--shadow);animation:slideIn 0.22s ease}
  @keyframes slideIn{from{opacity:0;transform:translateY(5px)}to{opacity:1;transform:none}}
  .job-card:hover{border-color:var(--border2);box-shadow:var(--shadow-md);transform:translateY(-1px)}
  .job-card.disabled{opacity:0.6}
  .job-status{width:8px;height:8px;border-radius:50%;flex-shrink:0}
  .job-status.on{background:var(--green);box-shadow:0 0 6px rgba(5,150,105,0.5)}
  .job-status.off{background:var(--text3)}
  .job-info{flex:1;min-width:0}
  .job-comment{font-size:13.5px;font-weight:600;color:var(--text);margin-bottom:5px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
  .job-command{font-family:'IBM Plex Mono',monospace;font-size:11px;color:#4f46e5;background:var(--accent-light);border:1px solid #c7d2fe;border-radius:5px;padding:2px 8px;display:inline-block;max-width:100%;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}
  .job-schedule{font-family:'IBM Plex Mono',monospace;font-size:11.5px;color:#7c3aed;background:#f5f3ff;border:1px solid #ddd6fe;padding:4px 10px;border-radius:20px;white-space:nowrap;flex-shrink:0;font-weight:500}
  .job-actions{display:flex;gap:6px;flex-shrink:0}

  .empty-state{text-align:center;padding:60px 20px;background:white;border:1.5px dashed var(--border2);border-radius:16px}
  .empty-icon{width:60px;height:60px;margin:0 auto 14px;background:var(--accent-light);border-radius:14px;display:flex;align-items:center;justify-content:center;font-size:26px}
  .empty-title{font-size:15px;font-weight:600;color:var(--text2);margin-bottom:5px}
  .empty-sub{font-size:13px;color:var(--text3)}

  /* MODAL */
  .modal-overlay{position:fixed;inset:0;background:rgba(15,23,42,0.55);backdrop-filter:blur(6px);z-index:1000;display:none;align-items:center;justify-content:center;padding:20px}
  .modal-overlay.open{display:flex}
  .modal{background:white;border:1.5px solid var(--border);border-radius:20px;width:100%;max-width:560px;max-height:92vh;overflow-y:auto;box-shadow:var(--shadow-lg);animation:modalIn 0.22s cubic-bezier(0.34,1.56,0.64,1)}
  @keyframes modalIn{from{opacity:0;transform:scale(0.94) translateY(10px)}to{opacity:1;transform:none}}
  .modal-header{padding:20px 22px 0;display:flex;align-items:center;justify-content:space-between;position:sticky;top:0;background:white;border-radius:20px 20px 0 0;z-index:1;padding-bottom:0}
  .modal-title-wrap{display:flex;align-items:center;gap:10px}
  .modal-title-icon{width:30px;height:30px;background:var(--accent-light);border-radius:8px;display:flex;align-items:center;justify-content:center;font-size:15px}
  .modal-title{font-size:15.5px;font-weight:700;color:var(--text)}
  .modal-close{background:none;border:none;cursor:pointer;color:var(--text3);font-size:17px;width:30px;height:30px;border-radius:7px;display:flex;align-items:center;justify-content:center;transition:all 0.12s}
  .modal-close:hover{background:var(--surface2);color:var(--text)}
  .modal-body{padding:16px 22px 20px}
  .modal-footer{padding:14px 22px;border-top:1.5px solid var(--border);display:flex;gap:10px;justify-content:space-between;align-items:center;background:var(--surface2);border-radius:0 0 20px 20px}

  /* WIZARD STEPS */
  .wizard-steps{display:flex;align-items:center;gap:0;margin:16px 0 20px;padding:0 2px}
  .wizard-step{display:flex;align-items:center;flex:1}
  .step-circle{width:26px;height:26px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-size:11px;font-weight:700;flex-shrink:0;transition:all 0.2s;border:2px solid var(--border)}
  .step-circle.done{background:var(--accent);border-color:var(--accent);color:white}
  .step-circle.active{background:var(--accent);border-color:var(--accent);color:white;box-shadow:0 0 0 3px rgba(79,70,229,0.2)}
  .step-circle.pending{background:white;border-color:var(--border2);color:var(--text3)}
  .step-label{font-size:11px;font-weight:600;margin-left:6px;color:var(--text3);white-space:nowrap}
  .step-label.active{color:var(--accent)}
  .step-label.done{color:var(--text2)}
  .step-connector{flex:1;height:2px;background:var(--border);margin:0 6px;transition:background 0.2s}
  .step-connector.done{background:var(--accent)}

  .wizard-page{display:none}
  .wizard-page.active{display:block}

  /* FORM */
  .form-group{margin-bottom:14px}
  .form-label{display:flex;align-items:center;gap:5px;font-size:11.5px;font-weight:700;color:var(--text2);margin-bottom:6px;text-transform:uppercase;letter-spacing:0.7px}
  .form-input,.form-select,.form-textarea{width:100%;background:white;border:1.5px solid var(--border);border-radius:9px;padding:9px 13px;color:var(--text);font-size:13.5px;font-family:'Plus Jakarta Sans',sans-serif;outline:none;transition:border-color 0.15s,box-shadow 0.15s}
  .form-input:focus,.form-select:focus,.form-textarea:focus{border-color:var(--accent);box-shadow:0 0 0 3px rgba(79,70,229,0.1)}
  .form-input::placeholder{color:var(--text3)}
  .form-select{cursor:pointer;appearance:none;background-image:url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='8' fill='none'%3E%3Cpath d='M1 1l5 5 5-5' stroke='%2394a3b8' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E");background-repeat:no-repeat;background-position:right 12px center;padding-right:34px}
  .form-textarea{min-height:100px;resize:vertical;font-family:'IBM Plex Mono',monospace;font-size:12px;line-height:1.6}
  .time-row{display:flex;gap:12px}
  .time-row .form-group{flex:1}
  .half-row{display:flex;gap:12px}
  .half-row .form-group{flex:1}

  /* Choice cards */
  .choice-grid{display:grid;grid-template-columns:repeat(3,1fr);gap:8px;margin-bottom:14px}
  .choice-card{background:white;border:1.5px solid var(--border);border-radius:10px;padding:12px 10px;text-align:center;cursor:pointer;transition:all 0.15s;user-select:none}
  .choice-card:hover{border-color:var(--border2);background:var(--surface2)}
  .choice-card.selected{border-color:var(--accent);background:var(--accent-light);color:var(--accent-text)}
  .choice-card-icon{font-size:20px;margin-bottom:4px}
  .choice-card-label{font-size:12px;font-weight:600}
  .choice-card-sub{font-size:10.5px;color:var(--text3);margin-top:2px}
  .choice-card.selected .choice-card-sub{color:var(--accent-text);opacity:0.75}

  /* Cron preview box */
  .cron-preview{background:var(--accent-light);border:1.5px solid #c7d2fe;border-radius:10px;padding:12px 16px;margin:16px 0;display:flex;align-items:center;gap:10px}
  .cron-preview-label{font-size:10px;font-weight:700;color:var(--accent-text);text-transform:uppercase;letter-spacing:0.7px;white-space:nowrap}
  .cron-preview-val{font-family:'IBM Plex Mono',monospace;font-size:13px;color:var(--accent-text);font-weight:500}
  .cron-human{font-size:12.5px;color:var(--text2);margin-top:6px;padding:8px 12px;background:var(--surface2);border-radius:7px;border:1px solid var(--border)}

  .info-box{font-size:12.5px;color:var(--text2);padding:10px 13px;background:var(--surface2);border:1px solid var(--border);border-radius:8px;margin-bottom:12px;line-height:1.5}
  .divider{height:1px;background:var(--border);margin:14px 0}
  .section-sub-label{font-size:11.5px;font-weight:700;text-transform:uppercase;letter-spacing:0.7px;color:var(--text3);margin-bottom:10px}

  .loading{text-align:center;padding:48px;color:var(--text3)}
  .spinner{width:26px;height:26px;border:2.5px solid var(--border);border-top-color:var(--accent);border-radius:50%;animation:spin 0.7s linear infinite;margin:0 auto 12px}
  @keyframes spin{to{transform:rotate(360deg)}}

  .toast{position:fixed;bottom:22px;right:22px;background:white;border:1.5px solid var(--border);border-radius:12px;padding:12px 18px;font-size:13.5px;font-weight:500;box-shadow:var(--shadow-lg);z-index:9999;display:flex;align-items:center;gap:9px;animation:toastIn 0.25s cubic-bezier(0.34,1.56,0.64,1);max-width:320px;color:var(--text)}
  @keyframes toastIn{from{opacity:0;transform:translateY(6px) scale(0.95)}to{opacity:1;transform:none}}
  .toast.success{border-left:3.5px solid var(--green)}
  .toast.error{border-left:3.5px solid var(--red)}
  ::-webkit-scrollbar{width:5px}
  ::-webkit-scrollbar-track{background:transparent}
  ::-webkit-scrollbar-thumb{background:var(--border2);border-radius:3px}
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
        <div class="logo-sub" data-i18n="sub"></div>
      </div>
    </div>
    <div class="nav">
      <div class="nav-label" data-i18n="nav_features"></div>
      <div class="nav-item active">
        <div class="nav-icon">⊞</div>
        <span data-i18n="nav_dashboard"></span>
      </div>
      <div class="nav-item" onclick="openJobModal()">
        <div class="nav-icon">＋</div>
        <span data-i18n="nav_add"></span>
      </div>
    </div>
    <div class="sidebar-footer">
      <div class="status-badge"><div class="status-dot"></div><span data-i18n="status_running"></span></div>
      <button class="lang-btn" id="lang-toggle" onclick="toggleLang()">EN</button>
    </div>
  </nav>

  <main class="main">
    <div class="page-header">
      <div>
        <div class="page-title" data-i18n="dashboard_title"></div>
        <div class="page-sub" data-i18n="dashboard_sub"></div>
      </div>
      <div style="display:flex;gap:10px;align-items:center;">
        <button class="btn btn-ghost" onclick="loadJobs()">↺ <span data-i18n="btn_refresh"></span></button>
        <button class="btn btn-primary" onclick="openJobModal()">＋ <span data-i18n="btn_add"></span></button>
      </div>
    </div>

    <div class="stats-row">
      <div class="stat-card total">
        <div class="stat-card-accent"></div>
        <div class="stat-label" data-i18n="stat_total_label"></div>
        <div class="stat-value" id="stat-total">—</div>
        <div class="stat-sub" data-i18n="stat_total_sub"></div>
      </div>
      <div class="stat-card active">
        <div class="stat-card-accent"></div>
        <div class="stat-label" data-i18n="stat_active_label"></div>
        <div class="stat-value" id="stat-active">—</div>
        <div class="stat-sub" data-i18n="stat_active_sub"></div>
      </div>
      <div class="stat-card disabled">
        <div class="stat-card-accent"></div>
        <div class="stat-label" data-i18n="stat_disabled_label"></div>
        <div class="stat-value" id="stat-disabled">—</div>
        <div class="stat-sub" data-i18n="stat_disabled_sub"></div>
      </div>
    </div>

    <div class="section-header">
      <div class="section-title" data-i18n="job_list_title"></div>
      <div id="last-refresh" style="font-size:11.5px;color:var(--text3)"></div>
    </div>
    <div id="job-list" class="job-list">
      <div class="loading"><div class="spinner"></div></div>
    </div>
  </main>
</div>

<!-- Job Modal (Add / Edit) -->
<div class="modal-overlay" id="job-modal">
  <div class="modal">
    <div class="modal-header" style="padding:20px 22px 16px;border-bottom:1.5px solid var(--border);">
      <div class="modal-title-wrap">
        <div class="modal-title-icon" id="modal-icon">⏰</div>
        <div class="modal-title" id="modal-title"></div>
      </div>
      <button class="modal-close" onclick="closeModal()">✕</button>
    </div>

    <!-- Wizard step indicators -->
    <div style="padding:0 22px">
      <div class="wizard-steps" id="wizard-steps">
        <div class="wizard-step" id="ws-0">
          <div class="step-circle active" id="sc-0">1</div>
          <div class="step-label active" id="sl-0" data-i18n="step_time"></div>
        </div>
        <div class="step-connector" id="conn-0"></div>
        <div class="wizard-step" id="ws-1">
          <div class="step-circle pending" id="sc-1">2</div>
          <div class="step-label" id="sl-1" data-i18n="step_freq"></div>
        </div>
        <div class="step-connector" id="conn-1"></div>
        <div class="wizard-step" id="ws-2">
          <div class="step-circle pending" id="sc-2">3</div>
          <div class="step-label" id="sl-2" data-i18n="step_cmd"></div>
        </div>
        <div class="step-connector" id="conn-2"></div>
        <div class="wizard-step" id="ws-3">
          <div class="step-circle pending" id="sc-3">4</div>
          <div class="step-label" id="sl-3" data-i18n="step_confirm"></div>
        </div>
      </div>
    </div>

    <div class="modal-body">

      <!-- Step 1: Time -->
      <div class="wizard-page active" id="page-0">
        <div class="section-sub-label" data-i18n="step1_heading"></div>
        <div class="time-row">
          <div class="form-group">
            <label class="form-label" data-i18n="lbl_hour"></label>
            <input type="number" class="form-input" id="f-hour" min="0" max="23" value="0" oninput="updatePreview()">
          </div>
          <div class="form-group">
            <label class="form-label" data-i18n="lbl_minute"></label>
            <input type="number" class="form-input" id="f-minute" min="0" max="59" value="0" oninput="updatePreview()">
          </div>
        </div>
      </div>

      <!-- Step 2: Frequency -->
      <div class="wizard-page" id="page-1">
        <div class="section-sub-label" data-i18n="step2_heading"></div>
        <!-- Day freq -->
        <div class="form-group">
          <label class="form-label" data-i18n="lbl_day_type"></label>
          <div class="choice-grid" id="day-type-grid">
            <div class="choice-card selected" data-val="daily" onclick="selectDayType('daily',this)">
              <div class="choice-card-icon">🔁</div>
              <div class="choice-card-label" data-i18n="opt_daily"></div>
              <div class="choice-card-sub" data-i18n="opt_daily_sub"></div>
            </div>
            <div class="choice-card" data-val="interval" onclick="selectDayType('interval',this)">
              <div class="choice-card-icon">📅</div>
              <div class="choice-card-label" data-i18n="opt_interval"></div>
              <div class="choice-card-sub" data-i18n="opt_interval_sub"></div>
            </div>
            <div class="choice-card" data-val="specific-day" onclick="selectDayType('specific-day',this)">
              <div class="choice-card-icon">📌</div>
              <div class="choice-card-label" data-i18n="opt_specific_day"></div>
              <div class="choice-card-sub" data-i18n="opt_specific_day_sub"></div>
            </div>
          </div>
        </div>
        <!-- N-day input -->
        <div id="day-interval-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_every_n_days"></label>
          <input type="number" class="form-input" id="f-days" min="1" max="365" value="2" oninput="updatePreview()">
        </div>
        <!-- Specific day of month -->
        <div id="day-specific-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_day_of_month"></label>
          <input type="number" class="form-input" id="f-monthday" min="1" max="31" value="1" oninput="updatePreview()">
        </div>

        <div class="divider"></div>

        <!-- Month freq -->
        <div class="form-group">
          <label class="form-label" data-i18n="lbl_month_type"></label>
          <div class="choice-grid" id="month-type-grid">
            <div class="choice-card selected" data-val="every-month" onclick="selectMonthType('every-month',this)">
              <div class="choice-card-icon">📆</div>
              <div class="choice-card-label" data-i18n="opt_every_month"></div>
              <div class="choice-card-sub" data-i18n="opt_every_month_sub"></div>
            </div>
            <div class="choice-card" data-val="month-interval" onclick="selectMonthType('month-interval',this)">
              <div class="choice-card-icon">🗓️</div>
              <div class="choice-card-label" data-i18n="opt_month_interval"></div>
              <div class="choice-card-sub" data-i18n="opt_month_interval_sub"></div>
            </div>
            <div class="choice-card" data-val="specific-month" onclick="selectMonthType('specific-month',this)">
              <div class="choice-card-icon">🎯</div>
              <div class="choice-card-label" data-i18n="opt_specific_month"></div>
              <div class="choice-card-sub" data-i18n="opt_specific_month_sub"></div>
            </div>
          </div>
        </div>
        <div id="month-interval-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_every_n_months"></label>
          <input type="number" class="form-input" id="f-month-interval" min="1" max="12" value="2" oninput="updatePreview()">
        </div>
        <div id="month-specific-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_which_month"></label>
          <select class="form-select" id="f-month" onchange="updatePreview()">
            <option value="1" data-i18n="m1"></option><option value="2" data-i18n="m2"></option>
            <option value="3" data-i18n="m3"></option><option value="4" data-i18n="m4"></option>
            <option value="5" data-i18n="m5"></option><option value="6" data-i18n="m6"></option>
            <option value="7" data-i18n="m7"></option><option value="8" data-i18n="m8"></option>
            <option value="9" data-i18n="m9"></option><option value="10" data-i18n="m10"></option>
            <option value="11" data-i18n="m11"></option><option value="12" data-i18n="m12"></option>
          </select>
        </div>

        <div class="divider"></div>

        <!-- Week freq -->
        <div class="form-group">
          <label class="form-label" data-i18n="lbl_week_type"></label>
          <div class="choice-grid" id="week-type-grid" style="grid-template-columns:repeat(3,1fr)">
            <div class="choice-card selected" data-val="every-week" onclick="selectWeekType('every-week',this)">
              <div class="choice-card-icon">📅</div>
              <div class="choice-card-label" data-i18n="opt_every_week"></div>
              <div class="choice-card-sub" data-i18n="opt_every_week_sub"></div>
            </div>
            <div class="choice-card" data-val="week-interval" onclick="selectWeekType('week-interval',this)">
              <div class="choice-card-icon">🔄</div>
              <div class="choice-card-label" data-i18n="opt_week_interval"></div>
              <div class="choice-card-sub" data-i18n="opt_week_interval_sub"></div>
            </div>
            <div class="choice-card" data-val="specific-weekday" onclick="selectWeekType('specific-weekday',this)">
              <div class="choice-card-icon">📍</div>
              <div class="choice-card-label" data-i18n="opt_specific_weekday"></div>
              <div class="choice-card-sub" data-i18n="opt_specific_weekday_sub"></div>
            </div>
          </div>
        </div>
        <div id="week-interval-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_every_n_weeks"></label>
          <input type="number" class="form-input" id="f-week-interval" min="1" max="52" value="2" oninput="updatePreview()">
        </div>
        <div id="week-specific-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_which_weekday"></label>
          <select class="form-select" id="f-weekday" onchange="updatePreview()">
            <option value="0" data-i18n="wd0"></option><option value="1" data-i18n="wd1"></option>
            <option value="2" data-i18n="wd2"></option><option value="3" data-i18n="wd3"></option>
            <option value="4" data-i18n="wd4"></option><option value="5" data-i18n="wd5"></option>
            <option value="6" data-i18n="wd6"></option>
          </select>
        </div>

        <div class="divider"></div>
        <!-- Custom -->
        <div class="form-group">
          <label style="display:flex;align-items:center;gap:8px;cursor:pointer;font-size:13px;font-weight:500;color:var(--text2)">
            <input type="checkbox" id="use-custom" onchange="toggleCustom()" style="width:14px;height:14px;accent-color:var(--accent)">
            <span data-i18n="use_custom_cron"></span>
          </label>
        </div>
        <div id="custom-cron-row" style="display:none" class="form-group">
          <label class="form-label" data-i18n="lbl_custom_cron"></label>
          <input type="text" class="form-input" id="f-custom" placeholder="* * * * *" style="font-family:'IBM Plex Mono',monospace;font-size:13px" oninput="updatePreview()">
          <div style="font-size:11px;color:var(--text3);margin-top:4px" data-i18n="custom_hint"></div>
        </div>
      </div>

      <!-- Step 3: Command -->
      <div class="wizard-page" id="page-2">
        <div class="form-group">
          <label class="form-label" data-i18n="lbl_cmd_type"></label>
          <select class="form-select" id="f-cmd-type" onchange="updateCmdType()">
            <option value="cmd" data-i18n="cmd_direct"></option>
            <option value="script-path" data-i18n="cmd_script_path"></option>
            <option value="script-content" data-i18n="cmd_script_content"></option>
          </select>
        </div>
        <div id="cmd-section-cmd">
          <div class="form-group">
            <label class="form-label" data-i18n="lbl_command"></label>
            <input type="text" class="form-input" id="f-command" style="font-family:'IBM Plex Mono',monospace;font-size:12.5px" data-i18n-placeholder="ph_command">
          </div>
        </div>
        <div id="cmd-section-script-path" style="display:none">
          <div class="form-group">
            <label class="form-label" data-i18n="lbl_script_path"></label>
            <input type="text" class="form-input" id="f-script-path" style="font-family:'IBM Plex Mono',monospace;font-size:12.5px" data-i18n-placeholder="ph_script_path">
          </div>
        </div>
        <div id="cmd-section-script-content" style="display:none">
          <div class="form-group">
            <label class="form-label" data-i18n="lbl_script_content"></label>
            <textarea class="form-textarea" id="f-script-content" data-i18n-placeholder="ph_script_content"></textarea>
            <div style="font-size:11.5px;color:var(--text3);margin-top:5px" data-i18n="script_note"></div>
          </div>
        </div>
        <div class="form-group">
          <label class="form-label"><span data-i18n="lbl_comment"></span> <span style="font-size:10px;font-weight:400;text-transform:none;letter-spacing:0;color:var(--text3)" data-i18n="optional"></span></label>
          <input type="text" class="form-input" id="f-comment" data-i18n-placeholder="ph_comment">
        </div>
      </div>

      <!-- Step 4: Confirm -->
      <div class="wizard-page" id="page-3">
        <div class="section-sub-label" data-i18n="step4_heading"></div>
        <div class="cron-preview">
          <div class="cron-preview-label">CRON</div>
          <div class="cron-preview-val" id="cron-preview-val">0 0 * * *</div>
        </div>
        <div class="cron-human" id="cron-human"></div>
        <div class="divider"></div>
        <div id="confirm-details" style="font-size:13px;color:var(--text2);line-height:1.7"></div>
      </div>

    </div><!-- modal-body -->

    <div class="modal-footer">
      <button class="btn btn-ghost" id="btn-back" onclick="wizardBack()" style="display:none" data-i18n="btn_back"></button>
      <div style="display:flex;gap:8px;margin-left:auto">
        <button class="btn btn-ghost" onclick="closeModal()" data-i18n="btn_cancel"></button>
        <button class="btn btn-primary" id="btn-next" onclick="wizardNext()" data-i18n="btn_next"></button>
      </div>
    </div>
  </div>
</div>

<script>
// ============================================================
// i18n
// ============================================================
const I18N = {
  zh: {
    sub:'定时任务管理器', nav_features:'功能', nav_dashboard:'仪表盘', nav_add:'添加任务',
    status_running:'服务运行中',
    dashboard_title:'仪表盘', dashboard_sub:'管理您的 Linux Crontab 定时任务',
    btn_refresh:'刷新', btn_add:'添加任务', btn_back:'上一步', btn_next:'下一步', btn_cancel:'取消', btn_save:'保存',
    stat_total_label:'全部任务', stat_total_sub:'已配置的定时任务',
    stat_active_label:'运行中', stat_active_sub:'当前启用的任务',
    stat_disabled_label:'已停用', stat_disabled_sub:'已禁用的任务',
    job_list_title:'任务列表',
    refreshed_at:'更新于 ', loading:'加载中...',
    empty_title:'暂无定时任务', empty_sub:'点击右上角「添加任务」开始',
    step_time:'时间', step_freq:'频率', step_cmd:'命令', step_confirm:'确认',
    step1_heading:'执行时间', step2_heading:'执行频率', step4_heading:'确认配置',
    lbl_hour:'小时 (0-23)', lbl_minute:'分钟 (0-59)',
    lbl_day_type:'每几天', lbl_month_type:'每几月', lbl_week_type:'每几周',
    opt_daily:'每天', opt_daily_sub:'每天执行',
    opt_interval:'每N天', opt_interval_sub:'指定间隔天数',
    opt_specific_day:'指定日期', opt_specific_day_sub:'每月第几天',
    opt_every_month:'每月', opt_every_month_sub:'每个月执行',
    opt_month_interval:'每N月', opt_month_interval_sub:'指定间隔月数',
    opt_specific_month:'指定月份', opt_specific_month_sub:'每年特定月份',
    opt_every_week:'每周', opt_every_week_sub:'每周执行',
    opt_week_interval:'每N周', opt_week_interval_sub:'指定间隔周数',
    opt_specific_weekday:'指定星期', opt_specific_weekday_sub:'每周特定天',
    lbl_every_n_days:'每隔几天', lbl_day_of_month:'每月第几天 (1-31)',
    lbl_every_n_months:'每隔几月', lbl_which_month:'哪个月份',
    lbl_every_n_weeks:'每隔几周', lbl_which_weekday:'星期几',
    use_custom_cron:'使用自定义 Cron 表达式（覆盖以上选项）',
    lbl_custom_cron:'Cron 表达式', custom_hint:'格式: 分 时 日 月 周，例如 0 2 * * 1',
    lbl_cmd_type:'命令类型', lbl_command:'执行命令', lbl_script_path:'脚本路径',
    lbl_script_content:'脚本内容', lbl_comment:'备注说明', optional:'(可选)',
    cmd_direct:'直接命令', cmd_script_path:'Shell 脚本路径', cmd_script_content:'编写 Shell 脚本',
    ph_command:'/usr/bin/python3 /home/user/script.py',
    ph_script_path:'/home/user/backup.sh',
    ph_script_content:'#!/bin/bash\necho Hello',
    ph_comment:'任务描述，便于识别...',
    script_note:'脚本将保存到服务器并自动执行',
    modal_add:'添加定时任务', modal_edit:'编辑定时任务',
    confirm_delete:'确认删除该任务？此操作不可撤销。',
    btn_enable:'启用', btn_disable:'停用', btn_edit:'编辑', btn_delete:'删除',
    toast_added:'任务添加成功！', toast_saved:'任务已保存', toast_deleted:'任务已删除',
    toast_err_cmd:'请填写命令', toast_err_cron:'Cron 格式错误，需要 5 个字段',
    toast_err_empty_cmd:'命令不能为空',
    m1:'1月',m2:'2月',m3:'3月',m4:'4月',m5:'5月',m6:'6月',
    m7:'7月',m8:'8月',m9:'9月',m10:'10月',m11:'11月',m12:'12月',
    wd0:'周日',wd1:'周一',wd2:'周二',wd3:'周三',wd4:'周四',wd5:'周五',wd6:'周六',
    human_daily:'每天 {H}:{M} 执行',
    human_interval:'每 {D} 天 {H}:{M} 执行',
    human_specific_day:'每月 {MD} 日 {H}:{M} 执行',
    human_every_month:'每月 {H}:{M}，每 {D} 天',
    human_custom:'自定义: {E}',
    confirm_cmd:'命令',confirm_comment:'备注',confirm_cron:'Cron 表达式',
  },
  en: {
    sub:'Cron Job Manager', nav_features:'Menu', nav_dashboard:'Dashboard', nav_add:'Add Job',
    status_running:'Service running',
    dashboard_title:'Dashboard', dashboard_sub:'Manage your Linux Crontab jobs',
    btn_refresh:'Refresh', btn_add:'Add Job', btn_back:'Back', btn_next:'Next', btn_cancel:'Cancel', btn_save:'Save',
    stat_total_label:'Total', stat_total_sub:'Configured jobs',
    stat_active_label:'Active', stat_active_sub:'Currently enabled',
    stat_disabled_label:'Disabled', stat_disabled_sub:'Paused jobs',
    job_list_title:'Job List',
    refreshed_at:'Updated at ', loading:'Loading...',
    empty_title:'No cron jobs yet', empty_sub:'Click "Add Job" to get started',
    step_time:'Time', step_freq:'Frequency', step_cmd:'Command', step_confirm:'Confirm',
    step1_heading:'Execution Time', step2_heading:'Frequency', step4_heading:'Review & Confirm',
    lbl_hour:'Hour (0-23)', lbl_minute:'Minute (0-59)',
    lbl_day_type:'Day Frequency', lbl_month_type:'Month Frequency', lbl_week_type:'Week Frequency',
    opt_daily:'Every Day', opt_daily_sub:'Run daily',
    opt_interval:'Every N Days', opt_interval_sub:'Set interval in days',
    opt_specific_day:'Specific Date', opt_specific_day_sub:'Day of month',
    opt_every_month:'Every Month', opt_every_month_sub:'Run every month',
    opt_month_interval:'Every N Months', opt_month_interval_sub:'Set month interval',
    opt_specific_month:'Specific Month', opt_specific_month_sub:'Certain months only',
    opt_every_week:'Every Week', opt_every_week_sub:'Run every week',
    opt_week_interval:'Every N Weeks', opt_week_interval_sub:'Set week interval',
    opt_specific_weekday:'Specific Weekday', opt_specific_weekday_sub:'Certain days only',
    lbl_every_n_days:'Every N days', lbl_day_of_month:'Day of month (1-31)',
    lbl_every_n_months:'Every N months', lbl_which_month:'Which month',
    lbl_every_n_weeks:'Every N weeks', lbl_which_weekday:'Day of week',
    use_custom_cron:'Use custom Cron expression (overrides above)',
    lbl_custom_cron:'Cron Expression', custom_hint:'Format: min hour day month weekday, e.g. 0 2 * * 1',
    lbl_cmd_type:'Command Type', lbl_command:'Command', lbl_script_path:'Script Path',
    lbl_script_content:'Script Content', lbl_comment:'Label', optional:'(optional)',
    cmd_direct:'Direct Command', cmd_script_path:'Shell Script Path', cmd_script_content:'Write Shell Script',
    ph_command:'/usr/bin/python3 /home/user/script.py',
    ph_script_path:'/home/user/backup.sh',
    ph_script_content:'#!/bin/bash\necho Hello',
    ph_comment:'Job description...',
    script_note:'Script will be saved on server and executed automatically',
    modal_add:'Add Cron Job', modal_edit:'Edit Cron Job',
    confirm_delete:'Delete this job? This cannot be undone.',
    btn_enable:'Enable', btn_disable:'Disable', btn_edit:'Edit', btn_delete:'Delete',
    toast_added:'Job added!', toast_saved:'Job saved', toast_deleted:'Job deleted',
    toast_err_cmd:'Please fill in the command', toast_err_cron:'Invalid Cron: need 5 fields',
    toast_err_empty_cmd:'Command cannot be empty',
    m1:'Jan',m2:'Feb',m3:'Mar',m4:'Apr',m5:'May',m6:'Jun',
    m7:'Jul',m8:'Aug',m9:'Sep',m10:'Oct',m11:'Nov',m12:'Dec',
    wd0:'Sun',wd1:'Mon',wd2:'Tue',wd3:'Wed',wd4:'Thu',wd5:'Fri',wd6:'Sat',
    human_daily:'Every day at {H}:{M}',
    human_interval:'Every {D} days at {H}:{M}',
    human_specific_day:'Day {MD} of every month at {H}:{M}',
    human_every_month:'Monthly at {H}:{M}',
    human_custom:'Custom: {E}',
    confirm_cmd:'Command',confirm_comment:'Label',confirm_cron:'Cron expression',
  }
};
let lang = 'zh';
function t(k){ return (I18N[lang][k]||I18N['zh'][k]||k); }
function applyI18n(){
  document.querySelectorAll('[data-i18n]').forEach(el=>{
    const k=el.getAttribute('data-i18n');
    el.textContent=t(k);
  });
  document.querySelectorAll('[data-i18n-placeholder]').forEach(el=>{
    el.placeholder=t(el.getAttribute('data-i18n-placeholder'));
  });
  document.getElementById('lang-toggle').textContent = lang==='zh'?'EN':'中文';
  updatePreview();
}
function toggleLang(){ lang = lang==='zh'?'en':'zh'; applyI18n(); }

// ============================================================
// Wizard state
// ============================================================
let currentStep = 0;
let editingId = null;
let dayType = 'daily';
let monthType = 'every-month';
let weekType = 'every-week';
let useCustom = false;
const TOTAL_STEPS = 4;

function openJobModal(editJob){
  editingId = editJob ? editJob.id : null;
  document.getElementById('modal-title').textContent = editJob ? t('modal_edit') : t('modal_add');
  document.getElementById('modal-icon').textContent = editJob ? '✏️' : '⏰';
  if(editJob){
    // Pre-fill form from existing job
    prefillEdit(editJob);
  } else {
    resetForm();
  }
  gotoStep(0);
  document.getElementById('job-modal').classList.add('open');
}

function closeModal(){
  document.getElementById('job-modal').classList.remove('open');
  editingId = null;
}
document.getElementById('job-modal').addEventListener('click',function(e){if(e.target===this)closeModal();});

function prefillEdit(job){
  resetForm();
  // Parse schedule back to form fields
  const parts = job.schedule.split(/\s+/);
  if(parts.length===5){
    document.getElementById('f-minute').value = parts[0]==='*'?'0':parts[0];
    document.getElementById('f-hour').value = parts[1]==='*'?'0':parts[1];
  }
  // Try to detect mode from schedule
  const [min,hour,dom,mon,dow] = parts;
  if(dom.startsWith('*/') && mon==='*' && dow==='*'){
    selectDayType('interval', document.querySelector('#day-type-grid [data-val="interval"]'));
    document.getElementById('f-days').value = dom.replace('*/','');
  } else if(dom!=='*' && mon==='*' && dow==='*'){
    selectDayType('specific-day', document.querySelector('#day-type-grid [data-val="specific-day"]'));
    document.getElementById('f-monthday').value = dom;
  }
  if(mon.startsWith('*/')){
    selectMonthType('month-interval', document.querySelector('#month-type-grid [data-val="month-interval"]'));
    document.getElementById('f-month-interval').value = mon.replace('*/','');
  } else if(mon!=='*'){
    selectMonthType('specific-month', document.querySelector('#month-type-grid [data-val="specific-month"]'));
    document.getElementById('f-month').value = mon;
  }
  if(dow!=='*'){
    selectWeekType('specific-weekday', document.querySelector('#week-type-grid [data-val="specific-weekday"]'));
    document.getElementById('f-weekday').value = dow;
  }
  // Try to detect if it's truly custom
  document.getElementById('f-command').value = job.command;
  document.getElementById('f-comment').value = job.comment||'';
  document.getElementById('f-cmd-type').value = 'cmd';
  updateCmdType();
  updatePreview();
}

// ============================================================
// Step navigation
// ============================================================
function gotoStep(n){
  currentStep = n;
  document.querySelectorAll('.wizard-page').forEach((p,i)=>p.classList.toggle('active',i===n));
  // Update step circles
  for(let i=0;i<TOTAL_STEPS;i++){
    const sc = document.getElementById('sc-'+i);
    const sl = document.getElementById('sl-'+i);
    sc.className = 'step-circle '+(i<n?'done':(i===n?'active':'pending'));
    sl.className = 'step-label '+(i<n?'done':(i===n?'active':''));
    if(i<TOTAL_STEPS-1){
      document.getElementById('conn-'+i).className='step-connector '+(i<n?'done':'');
    }
  }
  // Back button
  document.getElementById('btn-back').style.display = n>0?'':'none';
  // Next/Save button
  const btnNext = document.getElementById('btn-next');
  if(n===TOTAL_STEPS-1){
    btnNext.textContent = t('btn_save');
  } else {
    btnNext.textContent = t('btn_next');
  }
  if(n===TOTAL_STEPS-1) buildConfirmPage();
  updatePreview();
}

function wizardBack(){ if(currentStep>0) gotoStep(currentStep-1); }

function wizardNext(){
  if(currentStep===TOTAL_STEPS-1){ submitJob(); return; }
  // Validate current step
  if(currentStep===2){
    const ct = document.getElementById('f-cmd-type').value;
    let ok = false;
    if(ct==='cmd') ok = !!document.getElementById('f-command').value.trim();
    else if(ct==='script-path') ok = !!document.getElementById('f-script-path').value.trim();
    else if(ct==='script-content') ok = !!document.getElementById('f-script-content').value.trim();
    if(!ok){ showToast(t('toast_err_empty_cmd'),'error'); return; }
  }
  if(currentStep===1 && useCustom){
    const cron = document.getElementById('f-custom').value.trim();
    if(!cron||cron.split(/\s+/).length!==5){ showToast(t('toast_err_cron'),'error'); return; }
  }
  gotoStep(currentStep+1);
}

// ============================================================
// Day/Month/Week type selectors
// ============================================================
function selectDayType(val, el){
  dayType=val;
  document.querySelectorAll('#day-type-grid .choice-card').forEach(c=>c.classList.remove('selected'));
  if(el) el.classList.add('selected');
  document.getElementById('day-interval-row').style.display = val==='interval'?'':'none';
  document.getElementById('day-specific-row').style.display = val==='specific-day'?'':'none';
  updatePreview();
}
function selectMonthType(val, el){
  monthType=val;
  document.querySelectorAll('#month-type-grid .choice-card').forEach(c=>c.classList.remove('selected'));
  if(el) el.classList.add('selected');
  document.getElementById('month-interval-row').style.display = val==='month-interval'?'':'none';
  document.getElementById('month-specific-row').style.display = val==='specific-month'?'':'none';
  updatePreview();
}
function selectWeekType(val, el){
  weekType=val;
  document.querySelectorAll('#week-type-grid .choice-card').forEach(c=>c.classList.remove('selected'));
  if(el) el.classList.add('selected');
  document.getElementById('week-interval-row').style.display = val==='week-interval'?'':'none';
  document.getElementById('week-specific-row').style.display = val==='specific-weekday'?'':'none';
  updatePreview();
}
function toggleCustom(){
  useCustom = document.getElementById('use-custom').checked;
  document.getElementById('custom-cron-row').style.display = useCustom?'':'none';
  updatePreview();
}
function updateCmdType(){
  const t2 = document.getElementById('f-cmd-type').value;
  document.getElementById('cmd-section-cmd').style.display = t2==='cmd'?'':'none';
  document.getElementById('cmd-section-script-path').style.display = t2==='script-path'?'':'none';
  document.getElementById('cmd-section-script-content').style.display = t2==='script-content'?'':'none';
}

// ============================================================
// Cron builder
// ============================================================
function getCronExpr(){
  if(useCustom) return (document.getElementById('f-custom').value||'').trim()||'0 0 * * *';
  const h = document.getElementById('f-hour').value||'0';
  const m = document.getElementById('f-minute').value||'0';
  let dom='*', mon='*', dow='*';
  // Day
  if(dayType==='interval'){ const d=parseInt(document.getElementById('f-days').value)||2; dom='*/'+d; }
  else if(dayType==='specific-day'){ dom=document.getElementById('f-monthday').value||'1'; }
  // Month
  if(monthType==='month-interval'){ const mi=parseInt(document.getElementById('f-month-interval').value)||2; mon='*/'+mi; }
  else if(monthType==='specific-month'){ mon=document.getElementById('f-month').value||'1'; }
  // Weekday
  if(weekType==='week-interval'){ const wi=parseInt(document.getElementById('f-week-interval').value)||2; dow='*/'+wi; }
  else if(weekType==='specific-weekday'){ dow=document.getElementById('f-weekday').value||'0'; }
  return m+' '+h+' '+dom+' '+mon+' '+dow;
}

function getHumanCron(expr){
  const parts=expr.split(/\s+/);
  if(parts.length!==5) return expr;
  const[min,hour,dom,mon,dow]=parts;
  const pad=v=>String(v).padStart(2,'0');
  const H=pad(hour),M=pad(min);
  if(useCustom) return t('human_custom').replace('{E}',expr);
  if(dom==='*'&&mon==='*'&&dow==='*') return t('human_daily').replace('{H}',H).replace('{M}',M);
  if(dom.startsWith('*/')&&mon==='*'&&dow==='*'){
    return t('human_interval').replace('{D}',dom.replace('*/','')).replace('{H}',H).replace('{M}',M);
  }
  if(dom!=='*'&&mon==='*'&&dow==='*'){
    return t('human_specific_day').replace('{MD}',dom).replace('{H}',H).replace('{M}',M);
  }
  return t('human_custom').replace('{E}',expr);
}

function updatePreview(){
  const expr=getCronExpr();
  const el=document.getElementById('cron-preview-val');
  if(el) el.textContent=expr;
}

function buildConfirmPage(){
  const expr=getCronExpr();
  document.getElementById('cron-preview-val').textContent=expr;
  document.getElementById('cron-human').textContent=getHumanCron(expr);
  const ct=document.getElementById('f-cmd-type').value;
  let cmd='';
  if(ct==='cmd') cmd=document.getElementById('f-command').value;
  else if(ct==='script-path') cmd=document.getElementById('f-script-path').value;
  else cmd='['+t('cmd_script_content')+']';
  const comment=document.getElementById('f-comment').value;
  document.getElementById('confirm-details').innerHTML=
    '<b>'+t('confirm_cron')+':</b> <code style="font-family:IBM Plex Mono,monospace;font-size:12px;background:#f1f5f9;padding:1px 6px;border-radius:4px">'+escHtml(expr)+'</code><br>'+
    '<b>'+t('confirm_cmd')+':</b> <span style="font-family:IBM Plex Mono,monospace;font-size:12px;color:#4f46e5">'+escHtml(cmd)+'</span>'+(comment?'<br><b>'+t('confirm_comment')+':</b> '+escHtml(comment):'');
}

// ============================================================
// Submit
// ============================================================
async function submitJob(){
  const ct=document.getElementById('f-cmd-type').value;
  let command='',scriptPath='',scriptContent='';
  if(ct==='cmd') command=document.getElementById('f-command').value.trim();
  else if(ct==='script-path') scriptPath=document.getElementById('f-script-path').value.trim();
  else scriptContent=document.getElementById('f-script-content').value.trim();

  const body={
    mode: useCustom?'custom':(dayType==='daily'?'daily':(dayType==='interval'?'interval':'monthly')),
    days:document.getElementById('f-days').value||'2',
    weekday:document.getElementById('f-weekday').value||'0',
    monthDay:document.getElementById('f-monthday').value||'1',
    month:document.getElementById('f-month').value||'*',
    hour:document.getElementById('f-hour').value||'0',
    minute:document.getElementById('f-minute').value||'0',
    comment:document.getElementById('f-comment').value.trim(),
    customCron:getCronExpr(),
    command,scriptPath,scriptContent,
  };
  // Always send custom cron so backend uses it exactly
  body.mode='custom';

  const btn=document.getElementById('btn-next');
  btn.disabled=true;
  try {
    const url = editingId ? '/api/jobs/edit' : '/api/jobs/add';
    if(editingId) body.id=editingId;
    const res=await fetch(url,{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});
    const data=await res.json();
    if(data.success){
      showToast(editingId?t('toast_saved'):t('toast_added'),'success');
      closeModal(); loadJobs();
    } else showToast(data.message||'Error','error');
  } catch(e){ showToast(e.message,'error'); }
  finally{ btn.disabled=false; btn.textContent=t('btn_save'); }
}

// ============================================================
// Job list
// ============================================================
let jobs=[];
async function loadJobs(){
  document.getElementById('last-refresh').textContent='';
  try{
    const res=await fetch('/api/jobs');
    const data=await res.json();
    if(data.success){
      jobs=data.data||[];
      renderJobs(jobs); updateStats(jobs);
      document.getElementById('last-refresh').textContent=t('refreshed_at')+new Date().toLocaleTimeString(lang==='zh'?'zh-CN':'en-US');
    }
  } catch(e){ showToast(e.message,'error'); }
}

function renderJobs(jobs){
  const el=document.getElementById('job-list');
  if(!jobs||!jobs.length){
    el.innerHTML='<div class="empty-state"><div class="empty-icon">⏰</div><div class="empty-title">'+t('empty_title')+'</div><div class="empty-sub">'+t('empty_sub')+'</div></div>';
    return;
  }
  el.innerHTML=jobs.map(job=>{
    const isOn=job.enabled;
    const label=job.comment||(job.command.length>50?job.command.substring(0,50)+'...':job.command);
    const id=escAttr(job.id);
    return '<div class="job-card '+(isOn?'':'disabled')+'">'+
      '<div class="job-status '+(isOn?'on':'off')+'"></div>'+
      '<div class="job-info">'+
        '<div class="job-comment">'+escHtml(label)+'</div>'+
        '<span class="job-command">'+escHtml(job.command)+'</span>'+
      '</div>'+
      '<span class="job-schedule">'+escHtml(job.schedule)+'</span>'+
      '<div class="job-actions">'+
        '<button class="btn btn-ghost btn-sm" onclick="editJob(\''+id+'\')">✏️ '+t('btn_edit')+'</button>'+
        '<button class="btn '+(isOn?'btn-warning':'btn-ghost')+' btn-sm" onclick="toggleJob(\''+id+'\')">'+
          (isOn?'⏸ '+t('btn_disable'):'▶ '+t('btn_enable'))+'</button>'+
        '<button class="btn btn-danger btn-sm" onclick="deleteJob(\''+id+'\')">✕</button>'+
      '</div>'+
    '</div>';
  }).join('');
}

function updateStats(jobs){
  const total=jobs.length, active=jobs.filter(j=>j.enabled).length;
  document.getElementById('stat-total').textContent=total;
  document.getElementById('stat-active').textContent=active;
  document.getElementById('stat-disabled').textContent=total-active;
}

function editJob(id){
  const job=jobs.find(j=>j.id===id);
  if(job) openJobModal(job);
}

async function deleteJob(id){
  if(!confirm(t('confirm_delete'))) return;
  try{
    const res=await fetch('/api/jobs/delete',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id})});
    const data=await res.json();
    if(data.success){ showToast(t('toast_deleted'),'success'); loadJobs(); }
    else showToast(data.message,'error');
  } catch(e){ showToast(e.message,'error'); }
}

async function toggleJob(id){
  try{
    const res=await fetch('/api/jobs/toggle',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id})});
    const data=await res.json();
    if(data.success) loadJobs();
    else showToast(data.message,'error');
  } catch(e){ showToast(e.message,'error'); }
}

// ============================================================
// Helpers
// ============================================================
function resetForm(){
  document.getElementById('f-hour').value='0';
  document.getElementById('f-minute').value='0';
  document.getElementById('f-days').value='2';
  document.getElementById('f-monthday').value='1';
  document.getElementById('f-month-interval').value='2';
  document.getElementById('f-month').value='1';
  document.getElementById('f-week-interval').value='2';
  document.getElementById('f-weekday').value='0';
  document.getElementById('f-command').value='';
  document.getElementById('f-script-path').value='';
  document.getElementById('f-script-content').value='';
  document.getElementById('f-comment').value='';
  document.getElementById('f-custom').value='';
  document.getElementById('use-custom').checked=false;
  document.getElementById('f-cmd-type').value='cmd';
  useCustom=false;
  document.getElementById('custom-cron-row').style.display='none';
  updateCmdType();
  dayType='daily'; monthType='every-month'; weekType='every-week';
  selectDayType('daily', document.querySelector('#day-type-grid [data-val="daily"]'));
  selectMonthType('every-month', document.querySelector('#month-type-grid [data-val="every-month"]'));
  selectWeekType('every-week', document.querySelector('#week-type-grid [data-val="every-week"]'));
}

function escHtml(s){ return String(s||'').replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;'); }
function escAttr(s){ return String(s||'').replace(/['"\\]/g,c=>'\\'+c); }

let toastTimer;
function showToast(msg,type){
  const ex=document.querySelector('.toast'); if(ex) ex.remove();
  clearTimeout(toastTimer);
  const d=document.createElement('div');
  d.className='toast '+(type||'');
  d.innerHTML='<span>'+(type==='success'?'✓':'✕')+'</span>'+escHtml(msg);
  document.body.appendChild(d);
  toastTimer=setTimeout(()=>{ d.style.opacity='0'; d.style.transition='opacity 0.3s'; setTimeout(()=>d.remove(),300); },3000);
}

// ============================================================
// Init
// ============================================================
applyI18n();
loadJobs();
setInterval(loadJobs,30000);
</script>
</body>
</html>`
