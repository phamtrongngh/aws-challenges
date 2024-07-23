import { Body, Controller, Get, Inject, Post, Query } from '@nestjs/common';
import { Repository } from './repository/app.repository';

type PostMessageDto = {
  message: string;
};

@Controller()
export class AppController {
  constructor(
    @Inject('Repository')
    private readonly repo: Repository,
  ) {}

  @Post()
  async postMessage(@Body() data: PostMessageDto) {
    try {
      await this.repo.save(data.message);
      return {
        message: 'Message saved successfully',
      };
    } catch (e) {
      return {
        message: 'Error saving message',
      };
    }
  }

  @Get()
  async getMessages(@Query('limit') limit: number = 10) {
    try {
      const messages = await this.repo.get(limit);
      return {
        data: messages,
      };
    } catch (e) {
      return {
        message: 'Error fetching messages',
      };
    }
  }
}
