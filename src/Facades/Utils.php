<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Support\Common\Facade;
use Hyperf\Extra\Contract\UtilsServiceInterface;
use Ramsey\Uuid\Uuid;

/**
 * Class Utils
 * @package Hyperf\Support\Facades
 * @method static Uuid|\Ramsey\Uuid\UuidInterface uuid()
 */
final class Utils extends Facade
{
    protected static function getFacadeAccessor()
    {
        return UtilsServiceInterface::class;
    }
}