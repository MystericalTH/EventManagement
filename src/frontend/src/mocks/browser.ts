import { setupWorker } from 'msw/browser';
import { handlers } from './handlers';

// Configure the worker with the defined handlers
export const worker = setupWorker(...handlers);
