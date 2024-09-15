import React from 'react';
import ReactSlider from 'react-slider';
import './TimeSlider.css';

type TimeSliderProps = {
  minTime: number;
  maxTime: number;
  value: number;
  onChange: (value: number) => void;
};

const TimeSlider: React.FC<TimeSliderProps> = ({ minTime, maxTime, value, onChange }) => {
  return (
    <ReactSlider
      className="horizontal-slider"
      thumbClassName="thumb"
      trackClassName="track"
      min={minTime}
      max={maxTime}
      value={value}
      onChange={onChange}
      renderThumb={(props, state) => (
        <div {...props}>
          <div className="thumb-content">{new Date(state.valueNow).toLocaleString()}</div>
        </div>
      )}
    />
  );
};

export default TimeSlider;
