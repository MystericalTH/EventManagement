import Empty from './trash-empty.png';
import Filled from './trash-filled.png';

import type { ImageSource } from '$lib/types/actionButtonIcons';

let imgsrc: ImageSource = {
	idle: Empty,
	enter: Filled,
	click: Filled
};

export default imgsrc;
