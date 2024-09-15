import React, { useState } from 'react';
import './css/RegisterEventForm.css'; 

const RegisterEventForm: React.FC = () => {
  const [eventName, setEventName] = useState('');
  const [description, setDescription] = useState('');
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [venueImage, setVenueImage] = useState<File | null>(null); // 会場図の画像を管理するステート

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault(); // フォームのデフォルトの送信動作を防止

    // フォームデータの処理をここに追加します
    console.log('イベント名:', eventName);
    console.log('説明:', description);
    console.log('開始日時:', startDate);
    console.log('終了日時:', endDate);

    if (venueImage) {
      console.log('会場図:', venueImage.name);
    }

    // TODO: 実際の登録処理を追加
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setVenueImage(e.target.files[0]); // 最初のファイルをセット
    }
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
          type="datetime-local"
          value={startDate}
          onChange={(e) => setStartDate(e.target.value)}
        />
        
        <label htmlFor="endDate">終了日時</label>
        <input
          id="endDate"
          type="datetime-local"
          value={endDate}
          onChange={(e) => setEndDate(e.target.value)}
        />

		{/* バックエンドちゃんと書かないととても脆弱 */}
        <label htmlFor="venueImage">会場図の画像</label>
        <input
          id="venueImage"
          type="file"
          accept="image/*" 
          onChange={handleFileChange}
        />

        <button type="submit">Register</button>
      </form>
    </div>
  );
};

export default RegisterEventForm;
