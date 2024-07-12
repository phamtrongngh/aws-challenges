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
  postMessage(@Body() data: PostMessageDto) {
    this.repo.save(data.message);
  }

  @Get()
  async getMessages(@Query('limit') limit: number) {
    return this.repo.get(limit);
  }
}
