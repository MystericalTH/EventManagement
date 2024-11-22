import CrossDefault from './cross-gray.png';
import CrossRed from './cross-red.png';

import type { ImageSource } from '$lib/types';

let imgsrc: ImageSource = {
	idle: CrossDefault,
	enter: CrossRed,
	click: CrossRed
};

export default imgsrc;
