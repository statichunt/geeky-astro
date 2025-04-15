import React, { useEffect, useState, type JSX } from "react";

interface FontInfo {
  tag: string;
  fontSizePx: string;
  fontSizeRem: string;
}

const TextInfoList: React.FC = () => {
  const textElements = ["h1", "h2", "h3", "h4", "h5", "h6", "p"];

  const [fontInfo, setFontInfo] = useState<FontInfo[]>([]);

  useEffect(() => {
    const newFontInfo: FontInfo[] = textElements.map((tag) => {
      // Create a temporary element
      const tempElement = document.createElement(tag);
      tempElement.style.position = "absolute"; // Avoid affecting layout
      tempElement.style.visibility = "hidden";
      tempElement.textContent = `Temporary ${tag.toUpperCase()}`;

      // Append to the document body
      document.body.appendChild(tempElement);

      // Get computed styles
      const computedStyle = window.getComputedStyle(tempElement);
      const fontSizePx = `${Math.round(parseFloat(computedStyle.fontSize))}px`; // Rounded px value
      const fontSizeRem =
        (parseFloat(computedStyle.fontSize) / 16).toFixed(2) + "rem";

      // Remove the temporary element
      document.body.removeChild(tempElement);

      return {
        tag,
        fontSizePx,
        fontSizeRem,
      };
    });

    setFontInfo(newFontInfo);
  }, []);

  return (
    <div>
      {fontInfo.map(({ tag, fontSizePx, fontSizeRem }) => {
        const Tag = tag as keyof JSX.IntrinsicElements;
        return (
          <div key={tag} style={{ marginBottom: "1rem" }}>
            <Tag>{`${tag.toUpperCase()}: Font Size: ${fontSizePx} (${fontSizeRem})`}</Tag>
          </div>
        );
      })}
    </div>
  );
};

export default TextInfoList;
