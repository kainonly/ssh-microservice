<?php
declare(strict_types=1);

namespace Hyperf\Support\Common;

use Hyperf\Utils\ApplicationContext;
use Psr\Container\ContainerInterface;

abstract class RedisModel
{
    /**
     * @var ContainerInterface
     */
    protected $container;
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
    public function __construct(?\Redis $redis)
    {
        $this->container = ApplicationContext::getContainer();
        $this->redis = !empty($redis) ? $redis : $this->container->get(\Redis::class);
    }
}
