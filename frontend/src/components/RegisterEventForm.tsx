import React, { useState } from 'react';
import './css/RegisterEventForm.css'; 

const RegisterEventForm: React.FC = () => {
  const [eventName, setEventName] = useState('');
  const [description, setDescription] = useState('');
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault(); // フォームのデフォルトの送信動作を防止

    // フォームデータの処理をここに追加します
    console.log('イベント名:', eventName);
    console.log('説明:', description);
    console.log('開始日:', startDate);
    console.log('終了日:', endDate);

    // TODO: 実際の登録処理を追加
  };

  return (
    <div>
      <form className="register-event-form" onSubmit={handleSubmit}>
        <label htmlFor="eventName">イベント名</label>
        <input
          id="eventName"
          type="text"
          placeholder="イベント名"
          value={eventName}
          onChange={(e) => setEventName(e.target.value)}
        />
        
        <label htmlFor="description">説明</label>
        <input
          id="description"
          type="text"
          placeholder="説明"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        
        <label htmlFor="startDate">開始日時</label>
        <input
          id="startDate"
          type="datetime-local" // datetime-localで日付と時間を入力
          value={startDate}
          onChange={(e) => setStartDate(e.target.value)}
        />
        
        <label htmlFor="endDate">終了日時</label>
        <input
          id="endDate"
          type="datetime-local" // datetime-localで日付と時間を入力
          value={endDate}
          onChange={(e) => setEndDate(e.target.value)}
        />
        
        <button type="submit">Register</button>
      </form>
    </div>
  );
};

export default RegisterEventForm;
