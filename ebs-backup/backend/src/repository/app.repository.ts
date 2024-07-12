import { Injectable } from '@nestjs/common';
import { appendFile, readFile, mkdir, access } from 'fs/promises';
import * as path from 'path';

export interface Repository {
  save(msg: string): Promise<void>;
  get(limit: number): Promise<string[]>;
}

interface FileRepositoryOptions {
  filePath: string;
}

@Injectable()
export class FileRepository implements Repository {
  private filePath: string;

  constructor(private readonly options: FileRepositoryOptions) {
    this.filePath = options.filePath;
    this.createDataFile();
  }

  private async createDataFile() {
    try {
      await mkdir(path.dirname(this.filePath), { recursive: true });

      try {
        await access(this.filePath);
      } catch (err) {
        if (err.code === 'ENOENT') {
          await appendFile(this.filePath, '');
        } else {
          throw err;
        }
      }
    } catch (err) {
      console.error(err);
      throw err;
    }
  }

  async save(msg: string): Promise<void> {
    try {
      const data = `${new Date().toISOString()} - ${msg}\n`;
      await appendFile(this.filePath, data);
    } catch (error) {
      console.error('Failed to save message:', error);
      throw error;
    }
  }

  async get(limit?: number): Promise<string[]> {
    try {
      const content = await readFile(this.filePath, 'utf-8');
      const res = content.split('\n');
      res.splice(-1, 1);
      return res.slice(0, limit);
    } catch (error) {
      console.error('Failed to get messages:', error);
      throw error;
    }
  }
}
