import React from "react";

type MapProps = {
	src: string;
};

const Map: React.FC<MapProps> = ({ src }) => {
	return (
    <picture>
	    <source srcSet={src} type="image/webp" />
	    <img src={src} alt="会場地図" />
  	</picture>
  );
}

export default Map;