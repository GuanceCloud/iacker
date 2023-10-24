import React from 'react';
import TextTransition, { presets } from 'react-text-transition';

interface TextCarouselProps {
    texts: string[]
}

export default function TextCarousel(props: TextCarouselProps) {
  const [index, setIndex] = React.useState(0);

  React.useEffect(() => {
    const intervalId = setInterval(
      () => setIndex((index) => index + 1),
      3000, // every 3 seconds
    );
    return () => clearTimeout(intervalId);
  }, []);

  return (
    <TextTransition springConfig={presets.wobbly}>{props.texts[index % props.texts.length]}</TextTransition>
  );
};
