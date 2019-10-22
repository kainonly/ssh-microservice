<?php
declare(strict_types=1);

namespace Hyperf\Support\Common;

use Hyperf\Utils\ApplicationContext;

abstract class RedisModel
{
    /**
     * Model key
     * @var string $key
     */
    protected $key;

    /**
     * Redis Manager
     * @var  \Redis $redis
     */
    protected $redis;

    /**
     * RedisModel constructor.
     * @param \Redis $redis
     */
    public function __construct(\Redis $redis = null)
    {
        $this->redis = !empty($redis) ?
            $redis : ApplicationContext::getContainer()->get(\Redis::class);
    }
}
