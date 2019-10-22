<?php
declare(strict_types=1);

namespace Hyperf\Support\Common;

use Hyperf\Server\Exception\RuntimeException;
use Hyperf\Utils\ApplicationContext;

abstract class Facade
{
    /**
     * Create Facade Object
     * @return mixed
     * @throws \RuntimeException
     */
    protected static function createFacade()
    {
        $container = ApplicationContext::getContainer();
        $facadeClass = static::getFacadeAccessor();
        if (!$container->has($facadeClass)) {
            throw new RuntimeException('Injection dependency does not exist.');
        }
        return $container->get($facadeClass);
    }

    /**
     * Get the registered name of the component.
     * @return string
     * @throws \RuntimeException
     */
    protected static function getFacadeAccessor()
    {
        throw new RuntimeException('Facade does not implement getFacadeAccessor method.');
    }

    /**
     * @param $name
     * @param $arguments
     * @return mixed
     */
    public static function __callStatic($method, $arguments)
    {
        return call_user_func_array(
            [static::createFacade(), $method],
            $arguments
        );
    }
}