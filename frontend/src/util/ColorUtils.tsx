// 赤色に似ていないランダムな色を生成する関数
const generateRandomColor = () => {
  let color;
  do {
    color = `#${Math.floor(Math.random() * 16777215).toString(16).padStart(6, '0')}`;
  } while (/^#(FF|CC|99|66|33|00)0{2}/i.test(color)); // 赤色系を除外
  return color;
};

export { generateRandomColor };